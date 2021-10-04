package objects

import (
	"encoding/json"
	"net/http"
)

// Maximum listing
const MaxListLimit = 200

// Get request
type GetRequest struct {
	ID string `json:"id"`
}

// List request
type ListRequest struct {
	Limit int `json:"limit"`
	After int `json:"after"`
	Name  int `json:"name"`
}

// Create request
type CreateRequest struct {
	Event *Event `json:"event"`
}

type UpdateRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Website     string `json:"website"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type CancelRequest struct {
	ID string `json:"id"`
}

type RescheduleRequest struct {
	ID      string    `json:"id"`
	NewSlot *TimeSlot `json:"new_slot"`
}

type DeleteRequest struct {
	ID string `json:"id"`
}

// Response of every Event Request
type EventResponseWrapper struct {
	Event  *Event   `json:"event,omitempty`
	Events []*Event `json:"events,omitempty`
	Code   int      `json:"-"`
}

// EventResponseWrapper into JSON
func (e *EventResponseWrapper) JSON() []byte {
	if e == nil {
		return []byte("{}")
	}

	res, _ := json.Marshall(e)
	return res
}

// Status Code
func (e *EventResponseWrapper) StatusCode() int {
	if e == nil || e.Code == 0 {
		return http.StatusOK
	}

	return e.Code
}
