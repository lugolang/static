# static
lu static file serving middleware, based on fasthttp.FS.

## example
```go
package main

import (
	"log"
	
	"github.com/valyala/fasthttp"
	"github.com/vincentLiuxiang/lu"
    "github.com/lugolang/static"
)

func main() {
	app := lu.New()
	Static := static.New("/src/octopus-open/dist", "index.html")
	app.Get("/static", Static)
	server := &fasthttp.Server{
		Handler:       app.Handler,
		Concurrency:   1024 * 1024,
		Name:          "lu",
		MaxConnsPerIP: 50,
	}
	if err := server.ListenAndServe(":8080"); err != nil {
		log.Fatalf("error in lu server: %s", err)
	}
}

```