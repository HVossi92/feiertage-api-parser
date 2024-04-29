# feiertage-api-parser

This is a CLI parser to extract the dates per holiday for a date range using the feiertage-api service:

- https://feiertage-api.de/
- https://github.com/bundesAPI/feiertage-api

## Cli commands & flags

- -start the first year to fetch
- -end the last year to fetch (inclusive)
- `go run main.go -start 2019 -end 2030` (will fetch and parse all holidays from 2015 up to, and including, 2030 and save them to `./parsed-holidays.json`)

## JSON

The original feiertage-api provides the holidays for one year in the following format:

```json
{
  "BB": {
    "1. Weihnachtstag": {
      "datum": "2019-12-25",
      "hinweis": ""
    },
    "2. Weihnachtstag": {
      "datum": "2019-12-26",
      "hinweis": ""
    },
    "Christi Himmelfahrt": {
      "datum": "2019-05-30",
      "hinweis": ""
    },
    "etc": "..."
  },
  "BE": {
    "1. Weihnachtstag": {
      "datum": "2019-12-25",
      "hinweis": ""
    },
    "2. Weihnachtstag": {
      "datum": "2019-12-26",
      "hinweis": ""
    },
    "etc": "..."
  },
  "etc": "..."
}
```

This parser fetches the holidays for a range of years (e.g. 2015-2030) and converts them into an array of following format:

```json
[
  {
    "Name": "1. Weihnachtstag",
    "datum": "2015-12-25",
    "hinweis": ""
  },
  {
    "Name": "2. Weihnachtstag",
    "datum": "2015-12-26",
    "hinweis": ""
  },
  {
    "Name": "Christi Himmelfahrt",
    "datum": "2015-05-30",
    "hinweis": ""
  },
  {
    "Name": "Karfreitag",
    "datum": "2015-04-19",
    "hinweis": ""
  },
  {
    "etc": "..."
  },
  {
    "Name": "Buß- und Bettag",
    "datum": "2030-11-20",
    "hinweis": "Gemäß Art. 4 Nr. 3 des Bayerischen Feiertagsgesetzes[7] entfällt im gesamten Bundesland am Buß- und Bettag an allen Schulen der Unterricht."
  },
  {
    "Name": "Mariä Himmelfahrt",
    "datum": "2030-08-15",
    "hinweis": "Mariä Himmelfahrt ist in Bayern in von den derzeit 1704[8] (Zensus 2011, bis 2013"
  },
  {
    "Name": "Weltkindertag",
    "datum": "2030-09-20",
    "hinweis": ""
  }
]
```
