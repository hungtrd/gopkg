# swaggerui

A reusable Go package that serves Swagger UI with embedded static assets. Callers only need to provide the spec URL.

## Usage

```go
import "github.com/hungtrd/gopkg/swaggerui"

mux := http.NewServeMux()

// Mount at /swagger-ui/ (default)
swaggerui.Register(mux, "/openapi.yaml",
    swaggerui.WithTitle("My API"),
    swaggerui.WithDocExpansion("list"),
    swaggerui.WithTryItOut(),
)
```

This registers:

- `GET /swagger-ui` → 301 redirect to `/swagger-ui/`
- `GET /swagger-ui/` → Swagger UI

## Options

| Option                   | Default         | Description                                                |
| ------------------------ | --------------- | ---------------------------------------------------------- |
| `WithTitle(title)`       | `"Swagger UI"`  | HTML page title                                            |
| `WithBasePath(path)`     | `"/swagger-ui"` | Mount prefix                                               |
| `WithDocExpansion(mode)` | `"none"`        | Controls how API sections are expanded on load (see below) |
| `WithTryItOut()`         | disabled        | Enable try-it-out buttons                                  |

### WithDocExpansion modes

| Mode     | Behaviour                                                                                                                   |
| -------- | --------------------------------------------------------------------------------------------------------------------------- |
| `"none"` | All sections collapsed on load — fastest initial render, good for large specs                                               |
| `"list"` | Each tag group (e.g. "Users", "Orders") is expanded to show the list of endpoints, but each endpoint body remains collapsed |
| `"full"` | Every endpoint is fully expanded, showing request/response details immediately — convenient for small specs                 |

## Handler only (no mux registration)

```go
h := swaggerui.Handler("/openapi.yaml", swaggerui.WithTitle("My API"))
// h is an http.Handler; mount it however you like
```
