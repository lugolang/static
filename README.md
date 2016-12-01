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
	// fs.Root = "/static/file/path/"
	Static := static.New(fs)
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

the config of fs is totally same with fasthttp.FS

## fs.Root string

* The default value of fs.Root is ".", it means you should put your static file (html/css/js etc.) to the directory where your go program starts. 

* However you can custom it in your way.

## fs.IndexNames []string
* The default value is []string{"index.html"}
* It specify the index file. For example, when you access to http://xxxxxx:xxx/home/ , static will search the directory ```fs.Root + /home/``` to find the files in fs.IndexNames , and response to the client when find one.
* if no file is found , static will call next(nil) to pass the request to next non-error-middleware 