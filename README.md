# serve

A file serving middleware that works well (but not exclusively) with [rkusa/web](https://github.com/rkusa/web).

[![Build Status][travis]](https://travis-ci.org/rkusa/serve)
[![GoDoc][godoc]](https://godoc.org/github.com/rkusa/serve)

### Example

```go
app := web.New()
app.Use(serve.Dir("public"))
```

## License

[MIT](LICENSE)

[travis]: https://img.shields.io/travis/rkusa/serve.svg
[godoc]: http://img.shields.io/badge/godoc-reference-blue.svg
