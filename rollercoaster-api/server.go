package main

import (
	"encoding/json"
	"net/http"
)

type Coaster struct {
	Name         string `json:"name,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
	ID           string `json:"id,omitempty"`
	InPark       string `json:"in_park,omitempty"`
	Height       int    `json:"height,omitempty"`
}

type coasterHandlers struct {
	store map[string]Coaster
}

func (h *coasterHandlers) get(w http.ResponseWriter, r *http.Request) {

	coasters := make([]Coaster, len(h.store))

	i := 0
	for _, coaster := range h.store {
		coasters[i] = coaster
		i++
	}

	jsonBytes, err := json.Marshal(coasters)

	if err != nil {
		panic(err)
	}

	w.Write(jsonBytes)

}

func newCoasterHandlers() *coasterHandlers {
	return &coasterHandlers{
		store: map[string]Coaster{
			"id1": Coaster{
				Name:         "Fury",
				Height:       99,
				ID:           "id1",
				InPark:       "Carowinds",
				Manufacturer: "B+M"},
		},
	}
}

func main() {

	coasterHandlers := newCoasterHandlers()

	http.HandleFunc("/coasters", coasterHandlers.get)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
