package server_test

import (
	"github.com/matryer/is"
	"github.com/shawkyelshalawy/Daily_Brief/integrationtest"
	"net/http"
	"testing"
)

func TestServer_Start(t *testing.T) {
	integrationtest.SkipIfShort(t)

	t.Run("starts the server and listens for requests", func(t *testing.T) {
		is := is.New(t)

		cleanup := integrationtest.CreateServer()
		defer cleanup()

		resp, err := http.Get("http://localhost:8081/")
		is.NoErr(err)
		is.Equal(http.StatusOK, resp.StatusCode)
	})
}
