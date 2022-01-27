# HTTP API

## Analyze Text

POST: `/analysis`

This endpoint is used for analyzing input text for plagiarism based on given reference text.

**Body Fields:**

- `input_text`, String => input text that want to be analyzed.
- `ref_text`, String => text that will be used as reference for `input`.

**Example Request:**

```json
POST /analysis
Content-Type: application/json

{
    "input_text": "Why upon your first voyage as a passenger, did you yourself feel such a mystical vibration, when first told that you and your ship were now out of sight of land? Why did the old Persians hold the sea holy? But look! here come more crowds, pacing straight for the water, and seemingly bound for a dive. Strange! Nothing will content them but the extremest limit of the land; loitering under the shady lee of yonder warehouses will not suffice. No. They must get just as nigh the water as they possibly can without falling in. And there they stand—miles of them—leagues. Inlanders all, they come from lanes and alleys, streets and avenues—north, east, south, and west. Yet here they all unite. Tell me, does the magnetic virtue of the needles of the compasses of all those ships attract them thither? Go visit the Prairies in June, when for scores on scores of miles you wade knee-deep among tiger lilies what is the one charm wanting-water-there is not a drop of water there! Were Niagara but a cataract of sand, would you travel your thousand miles to see it?",
    "ref_text": "Once more. Say you are in the country; in some high land of lakes. Take almost any path you please, and ten to one it carries you down in a dale, and leaves you there by a pool in the stream. There is magic in it. Let the most absent-minded of men be plunged in his deepest reveries—stand that man on his legs, set his feet a-going, and he will infallibly lead you to water, if water there be in all that region. Should you ever be athirst in the great American desert, try this experiment, if your caravan happen to be supplied with a metaphysical professor. Yes, as every one knows, meditation and water are wedded for ever. But here is an artist. He desires to paint you the dreamiest, shadiest, quietest, most enchanting bit of romantic landscape in all the valley of the Saco. What is the chief element he employs? There stand his trees, each with a hollow trunk, as if a hermit and a crucifix were within; and here sleeps his meadow, and there sleep his cattle; and up from yonder cottage goes a sleepy smoke. Deep into distant woodlands winds a mazy way, reaching to overlapping spurs of mountains bathed in their hill-side blue. But though the picture lies thus tranced, and though this pine-tree shakes down its sighs like leaves upon this shepherd’s head, yet all were vain, unless the shepherd’s eye were fixed upon the magic stream before him. Go visit the Prairies in June, when for scores on scores of miles you wade knee-deep among Tiger-lilies—what is the one charm wanting?—Water—there is not a drop of water there! Were Niagara but a cataract of sand, would you travel your thousand miles to see it? Why did the poor poet of Tennessee, upon suddenly receiving two handfuls of silver, deliberate whether to buy him a coat, which he sadly needed, or invest his money in a pedestrian trip to Rockaway Beach? Why is almost every robust healthy boy with a robust healthy soul in him, at some time or other crazy to go to sea? Why upon your first voyage as a passenger, did you yourself feel such a mystical vibration, when first told that you and your ship were now out of sight of land? Why did the old Persians hold the sea holy? Why did the Greeks give it a separate deity, and own brother of Jove? Surely all this is not without meaning. And still deeper the meaning of that story of Narcissus, who because he could not grasp the tormenting, mild image he saw in the fountain, plunged into it and was drowned. But that same image, we ourselves see in all rivers and oceans. It is the image of the ungraspable phantom of life; and this is the key to it all."
}
```

**Success Response:**

```json
HTTP/1.1 200 OK
Content-Type: application/json

{
    "ok": true,
    "data": {
        "matches": [
            {
                "input": {
                    "text": "Why upon your first voyage as a passenger, did you yourself feel such a mystical vibration, when first told that you and your ship were now out of sight of land? Why did the old Persians hold the sea holy?",
                    "start_idx": 0,
                    "end_idx": 205
                },
                "ref": {
                    "text": "Why upon your first voyage as a passenger, did you yourself feel such a mystical vibration, when first told that you and your ship were now out of sight of land? Why did the old Persians hold the sea holy?",
                    "start_idx": 1939,
                    "end_idx": 2144
                }
            },
            {
                "input": {
                    "text": "Go visit the Prairies in June, when for scores on scores of miles you wade knee-deep among tiger lilies what is the one charm wanting-water-there is not a drop of water there! Were Niagara but a cataract of sand, would you travel your thousand miles to see it?",
                    "start_idx": 813,
                    "end_idx": 1073
                },
                "ref": {
                    "text": "Go visit the Prairies in June, when for scores on scores of miles you wade knee-deep among Tiger-lilies—what is the one charm wanting?—Water—there is not a drop of water there! Were Niagara but a cataract of sand, would you travel your thousand miles to see it?",
                    "start_idx": 1368,
                    "end_idx": 1629
                }
            }
        ]
    },
    "ts": 1643183609
}
```

**Error Responses:**

- Bad Request (`400`)

    ```json
    HTTP/1.1 400 Bad Request
    Content-Type: application/json

    {
        "ok": false,
        "err": "ERR_BAD_REQUEST",
        "msg": "missing `input_text`"
    }
    ```

- Internal Server Error (`500`)

    ```json
    HTTP/1.1 500 Internal Server Error
    Content-Type: application/json

    {
        "ok": false,
        "err": "ERR_INTERNAL_ERROR",
        "msg": "unable to connect to database"
    }
    ```