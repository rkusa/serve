package serve

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rkusa/web"
)

func TestDirFileFound(t *testing.T) {
	app := web.New()
	app.Use(Dir("."))

	secondCalled := false
	app.Use(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		secondCalled = true
		next(rw, r)
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
	app.Use(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		secondCalled = true
		next(rw, r)
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
