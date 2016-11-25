package static

import (
	"os"
	"github.com/valyala/fasthttp"
)

func New(path string, IndexNames string) func(ctx *fasthttp.RequestCtx, next func(error)) {
	if gopath := os.Getenv("GOPATH"); gopath == "" {
		panic(`System doesn't set env GOPATH, please set the env variable`)
	}
	fs := &fasthttp.FS{
		Root:     os.Getenv("GOPATH") + path,
		Compress: true,
	}
	if IndexNames != "" {
		fs.IndexNames = []string{IndexNames}
	}
	staticHandler := fs.NewRequestHandler()
	return func(ctx *fasthttp.RequestCtx, next func(error)) {
		staticHandler(ctx)
	}
}
