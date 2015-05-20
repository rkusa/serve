// A file serving middleware for [rkgo/web](https://github.com/rkgo/web)
//
//  app := web.New()
//  app.Use(serve.Dir("public"))
//
package serve

import (
	"net/http"

	"github.com/rkgo/web"
)

func Dir(dir string) web.Middleware {
	fs := http.Dir(dir)

	return func(ctx web.Context, next web.Next) {
		if ctx.Req().Method != "GET" && ctx.Req().Method != "HEAD" {
			next(ctx)
			return
		}

		name := ctx.Req().URL.Path
		file, err := fs.Open(name)
		if err != nil {
			next(ctx)
			return
		}
		defer file.Close()

		stat, err := file.Stat()
		if err != nil || stat.IsDir() {
			next(ctx)
			return
		}

		http.ServeContent(ctx, ctx.Req(), name, stat.ModTime(), file)
	}
}
