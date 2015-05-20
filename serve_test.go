package serve

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rkgo/web"
)

func TestDirFileFound(t *testing.T) {
	app := web.New()
	app.Use(Dir("."))

	secondCalled := false
	app.Use(func(ctx web.Context, next web.Next) {
		secondCalled = true
		next(ctx)
	})

	rec := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/serve.go", nil)
	if err != nil {
		t.Fatal(err)
	}

	app.ServeHTTP(rec, r)

	if rec.Code != http.StatusOK {
		t.Errorf("unexpected response, got: %d", rec.Code)
	}

	expectation, err := ioutil.ReadFile("serve.go")
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(rec.Body.Bytes(), expectation) {
		t.Error("files not equal")
	}

	if secondCalled != false {
		t.Error("second middleware called")
	}
}

func TestDirFileNotFound(t *testing.T) {
	app := web.New()
	app.Use(Dir("."))

	secondCalled := false
	app.Use(func(ctx web.Context, next web.Next) {
		secondCalled = true
		next(ctx)
	})

	rec := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/null", nil)
	if err != nil {
		t.Fatal(err)
	}

	app.ServeHTTP(rec, r)

	if rec.Code != http.StatusNotFound {
		t.Errorf("unexpected response, got: %d", rec.Code)
	}

	if secondCalled != true {
		t.Error("next middleware not called")
	}
}
