# Simple REST API


## Requirements
- [x] `GET /coasters` returns list of coasters
- [x] `GET /coasters/{id}` returns a specific coaster
- [x] `POST /coasters` accepts a new coaster
- [x] `POST /coasters` returns `415` if content is incorrect
- [x] `GET /coasters/random` returns a [random coaster](curl -v -L localhost:8080/coasters/random  | jq)


## Persistence
Temporary in memory storage

```
curl -v localhost:8080/coasters -X POST -d '{"name": "Taron2", "park": "Phantasialand", "height": 30, "manufacturer": "Intamin"}' -H "Content-Type: application/json"
```