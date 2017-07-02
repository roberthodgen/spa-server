# spa-server
> Go SPA (single page application) server handler

Get the package:

`go get github.com/roberthodgen/spa-server`

Use the package:

`import "github.com/roberthodgen/spa-server"`

For simple single page application serving using Go's `http` package.

## API

### `func SpaHandler(publicDir string, indexFile string) http.Handler`

Returns a request handler (`http.Handler`) that serves a single page application from a given directory (`publicDir`).

It falls back to a supplied index (`indexFile`) when either condition is true:

1. Request (file) path is not found
2. Request path is a directory


## Examples

Serve only a single page application:

```go
package main
import "net/http"
import spa "github.com/roberthodgen/spa-server"

func main() {
    http.ListenAndServe(":8080", spa.SpaHandler("public", "index.html"))
}
```

Serve a single page application along side other paths:

```go
package main
import "net/http"
import spa "github.com/roberthodgen/spa-server"

func main() {
    http.Handle("/app", spa.SpaHandler("webapp", "app.html"))
    // other handlers
    http.ListenAndServe(":8080", nil)
}
```
