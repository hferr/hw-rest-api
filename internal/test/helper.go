package test

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/hferr/hw-rest-api/internal/httpjson"
)

func Ptr[T any](v T) *T { return &v }

func DoHttpRequest(handler *httpjson.Handler, method, target string, body io.Reader) *http.Response {
	req := httptest.NewRequest(method, target, body)
	req.Header.Add("Content-Type", "application/json;charset=utf8")

	w := httptest.NewRecorder()

	handler.NewRouter().ServeHTTP(w, req)

	return w.Result()
}
