package ui

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"

	"github.com/ljgago/shortlink-go/ui/assets"
	"github.com/ljgago/shortlink-go/ui/pages"
)

func RegisterRoutes(r chi.Router) {
	r.Route("/assets/css", func(r chi.Router) {
		r.Handle("/*", http.StripPrefix("/assets", http.FileServerFS(assets.AssetsFS)))
	})

	r.Get("/", templ.Handler(pages.Home()).ServeHTTP)
}
