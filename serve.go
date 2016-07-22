// Package serve is a file serving middleware. It works well (but not
// exclusively) with [rkusa/web](https://github.com/rkusa/web).
//
//  app := web.New()
//  app.Use(serve.Dir("public"))
//
package serve

import (
	"net/http"
)

func Dir(dir string) func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	fs := http.Dir(dir)

	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if r.Method != "GET" && r.Method != "HEAD" {
			next(rw, r)
			return
		}

		name := r.URL.Path
		file, err := fs.Open(name)
		if err != nil {
			next(rw, r)
			return
		}
		defer file.Close()

		stat, err := file.Stat()
		if err != nil || stat.IsDir() {
			next(rw, r)
			return
		}

		http.ServeContent(rw, r, name, stat.ModTime(), file)
	}
}
