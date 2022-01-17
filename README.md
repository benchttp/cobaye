# cobaye

Local HTTP server exposing configurable endpoints to be benchmarked for testing purposes.


## Run

```sh
go run ./...
```

Note: default port is `9999`. It can be overrided with flag `-port`:

```sh
go run ./... -port 80
```
## Endpoints

The server exposes a single endpoint:

`GET /`

Optional query params:
- `delay` (integer): minimum duration in milliseconds before response
- `fib` (integer): nth element in fibonacci's suite to calculate before reponse
