package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/shawkyelshalawy/Daily_Brief/views"
	"net/http"
)

func FrontPage(mux chi.Router) {
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_ = views.FrontPage().Render(w)
	})
}
