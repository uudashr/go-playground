package outbox

import "time"

type Event struct {
	ID        string
	Name      string
	Body      []byte
	OccuredAt time.Time
}
