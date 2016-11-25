package static

import "github.com/valyala/fasthttp"
import "os"

func New(path string) func(ctx *fasthttp.RequestCtx, next func(error)) {
	fs := &fasthttp.FS{
		Root:       os.Getenv("GOPATH") + path,
		IndexNames: []string{"index.html"},
		Compress:   true,
	}
	staticHandler := fs.NewRequestHandler()
	return func(ctx *fasthttp.RequestCtx, next func(error)) {
		staticHandler(ctx)
	}
}
