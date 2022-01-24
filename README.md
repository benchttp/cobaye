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

## CLI

Typing the command `debug` in the CLI while the server is running will output
the count of received requests since it started.

## Endpoints

The server exposes two endpoints:

- `GET /`
  ```
  Response code: 200
  Response body: <empty>
  Optional query params:
  - `delay` (`time.Duration`): minimum duration before response
  - `fib` (`int`): nth element in fibonacci's suite to calculate before reponse
  ```

- `GET /debug`
  ```
  Response code: 200
  Response body: <received requests count> (raw text)
  ```

### Usage examples

- `http://localhost:9999?delay=250ms`
- `http://localhost:9999?fib=40`
- `http://localhost:9999?fib=40&delay=3s`
