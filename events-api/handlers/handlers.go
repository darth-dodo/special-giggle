package handlers

import (
	"net/http"
)

type EventHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	CancelEvent(w http.ResponseWriter, r *http.Request)
	RescheduleEvent(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type handler struct {
}

// NewEventHandler returns current interface EventHandler implementation

func NewEventHandler() EventHandler {
	return &handler{}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	panic("todo")
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	panic("todo")
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	panic("todo")
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	panic("todo")
}

func (h *handler) RescheduleEvent(w http.ResponseWriter, r *http.Request) {
	panic("todo")
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	panic("todo")
}

func (h *handler) CancelEvent(w http.ResponseWriter, r *http.Request) {
	panic("todo")
}
