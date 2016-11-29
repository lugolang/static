# static
[lu](https://github.com/vincentLiuxiang/lu) static file serving middleware, based on fasthttp.FS.

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
	fs := static.DefaultFS
	Static := static.New(*fs)
	app.Use("/static", Static)
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
