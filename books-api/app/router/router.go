package router

import (
	"github.com/darth-dodo/special-giggle/books-api/app/app"
	"github.com/go-chi/chi"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.MethodFunc("GET", "/", app.HandleIndex)

	return r
}