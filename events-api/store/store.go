package store

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/darth-dodo/special-giggle/events-api"
)

// EventStore is the database interface for storing the Events

type EventStore interface {
	Get(ctx context.Cotext, in *objects.GetRequest) (*objects.Event, error)
	List(ctx context.Cotext, in *objects.GetRequest) ([]*objects.Event, error)
	Update(ctx context.Cotext, in *objects.UpdateRequest) error
	Cancel(ctx context.Cotext, in *objects.CancelRequest) error
	Reschedule(ctx context.Context, in *objects.RescheduleRequest) error
	Delete(ctx context.Context, in *objects.DeleteRequest) error
}

func init() {
	rand.Seed(time.Now().UTC().Unix())
}

func GenerateUniqueID() string {
	word := []byte("0987654321")
	rand.Shuffle(len(word), func(i, j, int) {
		word[i], word[j] = word[j], word[i]
	})

	now := time.Now().UTC()
	return fmt.Sprintf("%010v-%010v-%s", now.Unix(), now.Nanosecond, string(word))
}
