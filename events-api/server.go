package main

import (
	"log"
	"net/http"

	"github.com/darth-dodo/special-giggle/events-api/handlers"
	"github.com/darth-dodo/special-giggle/events-api/store"
	"github.com/gorilla/mux"
)

// args used to run the server
type Args struct {
	// postgres connection string
	// e.g "postgres://user:password@localhost:5432/database?sslmode=disable"

	conn string
	port string
}

// run the server based on the given args
func Run(args Args) error {
	//router
	router := mux.NewRouter().PathPrefix("/api/v1/").Subrouter()

	database := store.NewPostgresEventStore(args.conn)
	controllers := handlers.NewEventHandler(database)

	RegisterAllRoutes(router, controllers)

	// start server
	log.Println("Starting the server at the port: ", args.port)
	return http.ListenAndServe(args.port, router)
}

// register all the routes of the API

func RegisterAllRoutes(router *mux.Router, controllers handlers.EventHandler) {
	// set the content type
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	// get events
	router.HandleFunc("/event", controllers.Get).Methods(http.MethodGet)

	// create events
	router.HandleFunc("/event", controllers.Create).Methods(http.MethodPost)

	// update events
	router.HandleFunc("/event", controllers.Update).Methods(http.MethodPut)

	//cancel events
	router.HandleFunc("/event/cancel", controllers.Cancel).Methods(http.MethodPatch)

	//delete events
	router.HandleFunc("/event/delete", controllers.Delete).Methods(http.MethodDelete)

	//reschedule events
	router.HandleFunc("/event/reschedule", controllers.Reschedule).Methods(http.MethodPatch)

	//list events
	router.HandleFunc("/events", controllers.List).Methods(http.MethodGet)

}
