package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Coaster struct {
	Name         string `json:"name,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
	ID           string `json:"id,omitempty"`
	InPark       string `json:"in_park,omitempty"`
	Height       int    `json:"height,omitempty"`
}

type coasterHandlers struct {
	sync.Mutex
	store map[string]Coaster
}

func (h *coasterHandlers) coasters(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return

	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
	}
}

func (h *coasterHandlers) get(w http.ResponseWriter, r *http.Request) {
	coasters := make([]Coaster, len(h.store))

	h.Lock()

	i := 0
	for _, coaster := range h.store {
		coasters[i] = coaster
		i++
	}

	h.Unlock()

	jsonBytes, err := json.Marshal(coasters)

	if err != nil {
		fmt.Printf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}

func (h *coasterHandlers) post(w http.ResponseWriter, r *http.Request) {

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		fmt.Printf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// check for header
	ct := r.Header.Get("content-type")

	if ct != "application/json" {

		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("need content type 'application/json' but got '%s'", ct)))
		return
	}

	//unmarshall bodyBytes to actual object

	var coaster Coaster
	err = json.Unmarshal(bodyBytes, &coaster)

	if err != nil {
		fmt.Printf(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	coaster.ID = fmt.Sprintf("%d", time.Now().UnixNano())

	h.Lock()
	h.store[coaster.ID] = coaster
	defer h.Unlock()
}

func (h *coasterHandlers) getCoaster(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.String(), "/")

	if len(parts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	h.Lock()
	coaster, ok := h.store[parts[2]]
	h.Unlock()

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonBytes, err := json.Marshal(coaster)

	if err != nil {
		fmt.Printf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
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

	http.HandleFunc("/coasters", coasterHandlers.coasters)
	http.HandleFunc("/coasters/", coasterHandlers.getCoaster)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}
