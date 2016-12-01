# static
[![Build Status](https://travis-ci.org/lugolang/static.svg?branch=master)](https://travis-ci.org/lugolang/static) [![Coverage Status](https://coveralls.io/repos/github/lugolang/static/badge.svg?branch=master)](https://coveralls.io/github/lugolang/static?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/lugolang/static)](https://goreportcard.com/report/github.com/lugolang/static)

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
