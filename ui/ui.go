package ui

import (
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"

	"github.com/ljgago/shortlink-go/ui/assets"
	"github.com/ljgago/shortlink-go/ui/pages"
)

func RegisterRoutes(r chi.Router) {
	isDevelopment := os.Getenv("GO_ENV") == "dev"

	r.Route("/assets/css", func(r chi.Router) {
		if isDevelopment {
				r.Handle("/*", http.StripPrefix("/assets", http.FileServer(http.Dir("ui/assets"))))
		} else {
				r.Handle("/*", http.StripPrefix("/assets", http.FileServerFS(assets.AssetsFS)))
		}
	})

	r.Get("/", templ.Handler(pages.Home()).ServeHTTP)
}
