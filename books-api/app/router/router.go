package router

import (
	"github.com/darth-dodo/special-giggle/books-api/app/app"
	"github.com/darth-dodo/special-giggle/books-api/app/requestlog"
	"github.com/go-chi/chi"
)

func New(a *app.App) *chi.Mux {
	l := a.Logger()

	r := chi.NewRouter()
	r.Method("GET", "/", requestlog.NewHandler(a.HandleIndex, l))

	return r
}
