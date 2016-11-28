package static

import (
	"errors"
	"os"

	"github.com/valyala/fasthttp"
)

var DefaultFS = &FS{
	&fasthttp.FS{
		Compress:        true,
		IndexNames:      []string{"index.html"},
		Root:            ".",
		AcceptByteRange: true,
	},
	true,
}

type FS struct {
	// Fasthttp *fasthttp.FS
	*fasthttp.FS
	// When something error occours,
	// FallThrough == true, next(nil)
	// FallThrough == false, next(err)
	FallThrough bool
}

func New(fs *FS) func(ctx *fasthttp.RequestCtx, next func(error)) {
	staticHandler := fs.NewRequestHandler()
	return func(ctx *fasthttp.RequestCtx, next func(error)) {
		m := string(ctx.Method())
		if m != "GET" && m != "HEAD" {
			if fs.FallThrough {
				next(nil)
				return
			}
			// method not allowed
			next(errors.New("405"))
			return
		}

		fileInfo, err := os.Stat(fs.Root + string(ctx.Path()))
		// if err != nil && os.IsNotExist(err) {
		if err != nil {
			if fs.FallThrough {
				next(nil)
				return
			}
			next(err)
			return
		}

		if !fileInfo.IsDir() {
			staticHandler(ctx)
			return
		}

		errPath := filterPath(fs.Root+string(ctx.Path())+"/", fs.IndexNames, next)

		if errPath != nil {
			if fs.FallThrough {
				next(nil)
				return
			}
			next(errPath)
			return
		}
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
