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

	Name        string `json:"name"`
	Description string `json:"description"`
	Website     string `json:"website"`
	Address     string `json:"address,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`

	// Slot duration
	Slot *TimeSlot `gorm:"embedded" json:"slot,omitempty"`

	// Status
	Status EventStatus `json:"status,omitempty"`

	// Meta information
	CreatedOn     time.Time `json:"created_on,omitempty"`
	UpdatedOn     time.Time `json:"updated_on,omitempty"`
	CancelledOn   time.Time `json:"cancelled_on,omitempty"`
	RescheduledOn time.Time `json:"rescheduled_on,omitempty"`
}
