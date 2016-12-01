package static

import (
	"strings"
	"testing"

	"github.com/vincentLiuxiang/lu"

	"os"

	"github.com/valyala/fasthttp"
)

func Test_Method(t *testing.T) {
	fi, _ := os.Create("index.html")
	fi.WriteString(`<!DOCTYPE html>
  <html lang="en">
  <head>
    <meta charset="UTF-8">
    <title>Document</title>
  </head>
  <body>
    <h2>hello static</h2>
  </body>
  </html>`)
	fi.Close()
	app := lu.New()
	fs := DefaultFS
	Static := New(*fs)
	app.Use("/", Static)

	postFallThroughMw := false
	postFallThroughErr := false

	app.Use("/", func(ctx *fasthttp.RequestCtx, next func(error)) {
		postFallThroughMw = true
		postFallThroughErr = false
	})

	go app.Listen(":3000")

	postFallThroughMw = false
	postFallThroughErr = false
	code, body, _ := fasthttp.Get(nil, "http://localhost:3000/")
	if code == 200 && strings.Contains(string(body), "hello static") {
		t.Log("OK")
	} else {
		t.Error("ERROR")
	}

	postFallThroughMw = false
	postFallThroughErr = true
	code, body, _ = fasthttp.Get(nil, "http://localhost:3000/test/xxx")
	if code == 200 && postFallThroughMw && !postFallThroughErr {
		t.Log("OK")
	} else {
		t.Error("ERROR")
	}

	postFallThroughMw = false
	postFallThroughErr = true
	code, _, _ = fasthttp.Get(nil, "http://localhost:3000/test/dist")
	if code == 200 && postFallThroughMw && !postFallThroughErr {
		t.Log("OK")
	} else {
		t.Error("ERROR")
	}
	os.Remove("index.html")
}

func Test_FallThrough1(t *testing.T) {
	os.Mkdir("test", os.ModePerm)
	os.Mkdir("test/dist", os.ModePerm)
	fi, _ := os.Create("test/index.html")
	fi.WriteString(`<!DOCTYPE html>
	<html lang="en">
	<head>
	  <meta charset="UTF-8">
	  <title>Document</title>
	</head>
	<body>
	  <h2>hello test</h2>
	</body>
	</html>`)
	fi.Close()

	fi, _ = os.Create("index.html")
	fi.WriteString(`<!DOCTYPE html>
	<html lang="en">
	<head>
	  <meta charset="UTF-8">
	  <title>Document</title>
	</head>
	<body>
	  <h2>hello static</h2>
	</body>
	</html>`)
	fi.Close()

	app := lu.New()
	fs := DefaultFS
	Static := New(*fs)
	app.Use("/", Static)

	var postFallThroughMw, postFallThroughErr bool
	app.Use("/", func(ctx *fasthttp.RequestCtx, next func(error)) {
		postFallThroughMw = true
		postFallThroughErr = false
	})

	go app.Listen(":3010")

	// method
	postFallThroughMw = false
	postFallThroughErr = true
	code, _, _ := fasthttp.Post(nil, "http://localhost:3010/", nil)
	if code == 200 && postFallThroughMw && !postFallThroughErr {
		t.Log("OK")
	} else {
		t.Error("ERROR")
	}

	postFallThroughMw = false
	postFallThroughErr = true
	code, _, _ = fasthttp.Get(nil, "http://localhost:3010/test/xxx")
	if code == 200 && postFallThroughMw && !postFallThroughErr {
		t.Log("OK")
	} else {
		t.Error("ERROR")
	}

	postFallThroughMw = false
	postFallThroughErr = true
	co, body, _ := fasthttp.Get(nil, "http://localhost:3010/test/index.html")
	if co == 200 && strings.Contains(string(body), "hello test") {
		t.Log("OK")
	} else {
		t.Error("ERROR")
	}

	postFallThroughMw = false
	postFallThroughErr = true
	c, _, _ := fasthttp.Get(nil, "http://localhost:3010/test/dist")
	if c == 200 && postFallThroughMw && !postFallThroughErr {
		t.Log("OK")
	} else {
		t.Error("ERROR")
	}
	os.RemoveAll("test")
	os.Remove("index.html")
}

func Test_FallThrough2(t *testing.T) {
	os.Mkdir("test", os.ModePerm)
	os.Mkdir("test/dist", os.ModePerm)
	fi, _ := os.Create("test/index.html")
	fi.WriteString(`<!DOCTYPE html>
  <html lang="en">
  <head>
    <meta charset="UTF-8">
    <title>Document</title>
  </head>
  <body>
    <h2>hello test</h2>
  </body>
  </html>`)

	fi.Close()
	app := lu.New()
	fs := DefaultFS
	Static := New(*fs)
	app.Use("/", Static)

	var postFallThroughMw, postFallThroughErr bool
	app.Use("/", func(ctx *fasthttp.RequestCtx, next func(error)) {
		postFallThroughMw = true
		postFallThroughErr = false
		ctx.SetStatusCode(404)
	})

	go app.Listen(":3011")

	// method
	postFallThroughMw = false
	postFallThroughErr = true

	code, _, _ := fasthttp.Post(nil, "http://localhost:3011/", nil)
	if code == 404 && postFallThroughMw && !postFallThroughErr {
		t.Log("OK")
	} else {
		t.Error("ERROR")
	}

	code, _, _ = fasthttp.Get(nil, "http://localhost:3011/test/dist")
	if code == 404 {
		t.Log("OK")
	} else {
		t.Error("ERROR")
	}
	os.RemoveAll("test")
}

func Test_FallThrough3(t *testing.T) {
	os.Mkdir("test", os.ModePerm)
	os.Mkdir("test/dist", os.ModePerm)
	fi, _ := os.Create("test/index.html")
	fi.WriteString(`<!DOCTYPE html>
  <html lang="en">
  <head>
    <meta charset="UTF-8">
    <title>Document</title>
  </head>
  <body>
    <h2>hello test</h2>
  </body>
  </html>`)

	fi.Close()
	app := lu.New()
	fs := DefaultFS
	fs.IndexNames = nil
	Static := New(*fs)
	app.Use("/", Static)

	var postFallThroughMw, postFallThroughErr bool
	app.Use("/", func(ctx *fasthttp.RequestCtx, next func(error)) {
		postFallThroughMw = true
		postFallThroughErr = false
		ctx.SetStatusCode(404)
	})

	go app.Listen(":3012")

	postFallThroughMw = false
	postFallThroughErr = true
	code, _, _ := fasthttp.Get(nil, "http://localhost:3012/test/")
	if code == 404 && postFallThroughMw && !postFallThroughErr {
		t.Log("OK")
	} else {
		t.Error("ERROR")
	}
	os.RemoveAll("test")
}
