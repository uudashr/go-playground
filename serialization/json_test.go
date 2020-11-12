package serialization_test

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/uudashr/go-playground/serialization"
)

func TestDeserialize(t *testing.T) {
	events := []interface{}{
		RegisterProductRequested{
			RequestID: "c69969a7-24f4-4ffc-877a-290eea63b3b3",
			ID:        "wood-spoon",
			Name:      "Wood Spoon",
		},
		AddItemToCartRequested{
			RequestID: "29dcb0e6-d13d-4587-8392-e389c0d2cc38",
			CartID:    "john-appleseed",
			ItemID:    "wood-spoon",
			Quantity:  12,
		},
	}

	var dez serialization.Deserializer

	dez.Register(RegisterProductRequested{})
	dez.Register(AddItemToCartRequested{})

	for i, event := range events {
		eventLog, err := LogEvent(event)
		if err != nil {
			t.Fatalf("i=%d LogEvent err: %v", i, err)
		}

		v, err := dez.Deserialize(eventLog.Name, eventLog.Body)
		if err != nil {
			t.Fatalf("i=%d Deserialize err: %v", i, err)
		}

		if got, want := v, event; !reflect.DeepEqual(got, want) {
			t.Fatalf("i=%d event got: %+v, want: %+v", i, got, want)
		}
	}
}

func LogEvent(event interface{}) (*EventLog, error) {
	b, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	return &EventLog{
		Name:      eventName(event),
		Body:      b,
		Timestamp: time.Now(),
	}, nil
}

func eventName(e interface{}) string {
	return reflect.TypeOf(e).Name()
}

type EventLog struct {
	Name      string    `json:"name"`
	Body      []byte    `json:"body"`
	Timestamp time.Time `json:"timestamp"`
}

type RegisterProductRequested struct {
	RequestID string
	ID        string
	Name      string
}

type AddItemToCartRequested struct {
	RequestID string
	CartID    string
	ItemID    string
	Quantity  int
}
