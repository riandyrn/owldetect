package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// define handlers
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/analysis", func(w http.ResponseWriter, r *http.Request) {
		// check http method
		if r.Method != http.MethodPost {
			WriteAPIResp(w, NewErrorResp(NewErrMethodNotAllowed()))
			return
		}
		// parse request body
		var reqBody analyzeReqBody
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			WriteAPIResp(w, NewErrorResp(NewErrBadRequest(err.Error())))
			return
		}
		// validate request body
		err = reqBody.Validate()
		if err != nil {
			WriteAPIResp(w, NewErrorResp(err))
			return
		}
		// do analysis
		matches := doAnalysis(reqBody.InputText, reqBody.RefText)
		// output success response
		WriteAPIResp(w, NewSuccessResp(map[string]interface{}{
			"matches": matches,
		}))
	})
	// define port, we need to set it as env for Heroku deployment
	port := os.Getenv("PORT")
	if port == "" {
		port = "9056"
	}
	// run server
	log.Printf("server is listening on :%v", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("unable to run server due: %v", err)
	}
}

func doAnalysis(input, ref string) []match {
	idx := strings.Index(ref, input)
	if idx == -1 {
		return nil
	}
	return []match{
		{
			Input: matchDetails{
				Text:     input,
				StartIdx: 0,
				EndIdx:   len(input) - 1,
			},
			Reference: matchDetails{
				Text:     ref[idx : idx+len(input)],
				StartIdx: idx,
				EndIdx:   idx + len(input) - 1,
			},
		},
	}
}
