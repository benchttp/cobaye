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
- `delay` (`time.Duration`): minimum duration before response
- `fib` (`int`): nth element in fibonacci's suite to calculate before reponse

Examples:
- `http://localhost:9999?delay=250ms`
- `http://localhost:9999?fib=40`
- `http://localhost:9999?fib=40&delay=3s`
