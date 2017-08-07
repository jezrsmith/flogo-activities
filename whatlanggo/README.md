# whatlanggo
This activity provides language and script detection given a text string.

Makes use of https://github.com/abadojack/whatlanggo which isa derivation of Franc (https://github.com/wooorm/franc)


## Installation

```bash
flogo add activity github.com/TIBCOSoftware/flogo-contrib/activity/whatlanggo
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "phrase",
      "type": "string",
      "required": true
    }
  ],
  "outputs": [
    {
      "name": "language",
      "type": "string"
    },
    {
      "name": "script",
      "type": "string"
    }
  ]
}
```
## Input
| Input     | Description    |
|:------------|:---------------|
| phrase | The phrase for which we want to identify the language |         

## Output
| Output     | Description    |
|:------------|:---------------|
| language | The detected language code |
| script | The detected language script |


## Example
Detect language of the phrase: NOTRE SÉLECTION D'OFFRES DE LA SEMAINE

```json
{
  "id": 3,
  "type": 1,
  "activityType": "whatlanggo",
  "name": "What language",
  "attributes": [
    { "name": "phrase", "value": "NOTRE SÉLECTION D'OFFRES DE LA SEMAINE" }
  ]
}
```