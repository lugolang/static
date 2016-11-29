package static

import (
	"errors"
	"os"

	"github.com/valyala/fasthttp"
)

var DefaultFS = &fasthttp.FS{
	Compress:        true,
	IndexNames:      []string{"index.html"},
	Root:            ".",
	AcceptByteRange: true,
}

// fs is a value which references to a *fasthttp.FS pointer.
// The reason why New accepts a fasthttp.FS value but not a *fasthttp.FS pointer
// is that after we called New() and return a lu Middleware, we don't
// hope anyone can change any property of fs(fasthttp.FS).
// However, if New accepts a *fasthttp.FS pointer, we can change properties
// of fasthttp.FS even after we called New().
//
// if the http request method is not 'GET' or 'HEAD', the static will call next(nil)
//
// if the http request file does not exist or the directory does not have the IndexNames file.
// the static will call next(nil)
//
// or, fasthttp.FS will handle the request file or the IndexNames file in the request directory
func New(fs fasthttp.FS) func(ctx *fasthttp.RequestCtx, next func(error)) {
	staticHandler := fs.NewRequestHandler()
	return func(ctx *fasthttp.RequestCtx, next func(error)) {
		m := string(ctx.Method())
		if m != "GET" && m != "HEAD" {
			next(nil)
			return
		}

		path := string(ctx.Path())

		fileInfo, err := os.Stat(fs.Root + path)
		// if err != nil && os.IsNotExist(err) {
		if err != nil {
			next(nil)
			return
		}

		// An exist file
		// fasthttp.FS handle it
		if !fileInfo.IsDir() {
			staticHandler(ctx)
			return
		}

		if len(fs.IndexNames) == 0 {
			next(nil)
			return
		}

		errPath := filterPath(fs.Root+path+"/", fs.IndexNames, next)
		if errPath != nil {
			next(nil)
			return
		}

		// An exist directory has IndexNames file
		staticHandler(ctx)
	}
}

func filterPath(path string, index []string, next func(error)) error {
	for _, v := range index {
		_, err := os.Stat(path + v)
		if err == nil {
			return nil
		}
	}
	return errors.New("Not found index page in directory: " + path)
}
