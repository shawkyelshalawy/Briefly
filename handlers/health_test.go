package handlers_test

import (
	"github.com/matryer/is"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/shawkyelshalawy/Daily_Brief/handlers"
)

func TestHealth(t *testing.T) {
	t.Run("retruns 200", func(t *testing.T) {
		is := is.New(t)
		mux := chi.NewRouter()
		handlers.Health(mux)
		code, _, _ := makeGetRequest(mux, "/health")
		is.Equal(http.StatusOK, code)
	})
}
func makeGetRequest(handler http.Handler, target string) (int, http.Header, string) {
	req := httptest.NewRequest(http.MethodGet, target, nil)
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)
	result := res.Result()
	bodyBytes, err := io.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}
	return result.StatusCode, result.Header, string(bodyBytes)
}
