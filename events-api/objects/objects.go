package objects

import "time"

// Defines the staus of the event
type EventStatus string

// Default event status
const (
	Original    EventStatus = "original"
	Cancelled   EventStatus = "cancelled"
	Rescheduled EventStatus = "rescheduled"
)

// time slot for event
type TimeSlot struct {
	StartTime time.Time `json:"start_time,omitempty"`
	EndTime   time.Time `json:"end_time,omitempty"`
}

// Event object for the API
type Event struct {
	ID string `gorm:"primary_key" json:"id,omitempty"`

	Name        string `json:"end_time,omitempty"`
	Description string `json:"end_time,omitempty"`
	Website     string `json:"end_time,omitempty"`
	Address     string `json:"end_time,omitempty"`
	PhoneNumber string `json:"end_time,omitempty"`

	// Slot duration
	Slot *TimeSlot `gorm:"embedded" json:"slot,omitempty"`

	// Status
	Status EventStatus `json:"status,omitempty"`

	// Meta information
	CreatedOn   time.Time `json:"created_on,omitempty"`
	UpdatedOn   time.Time `json:"updated_on,omitempty"`
	CancelledOn time.Time `json:"cancelled_on,omitempty"`
}
