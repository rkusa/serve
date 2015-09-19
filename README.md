# serve

A file serving middleware for [rkgo/web](https://github.com/rkgo/web)

[![Build Status][drone]](https://ci.rkusa.st/rkgo/serve)
[![GoDoc][godoc]](https://godoc.org/github.com/rkgo/serve)

### Example

```go
app := web.New()
app.Use(serve.Dir("public"))
```

[drone]: http://ci.rkusa.st/api/badges/rkgo/serve/status.svg?style=flat-square
[godoc]: http://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square