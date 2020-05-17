package rabbitmq_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/streadway/amqp"
)

func TestPublish(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	body := fmt.Sprintf("Hello World %d", rand.Int())

	// receiver
	rcvdone := make(chan struct{})
	rcvbody := make(chan string)
	{
		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		if err != nil {
			t.Error("Fail to dial:", err)
			return
		}

		defer func() {
			err = conn.Close()
			if err != nil {
				t.Log("Fail on close:", err)
			}
		}()

		ch, err := conn.Channel()
		if err != nil {
			t.Error("Fail to open channel:", err)
			return
		}

		q, err := ch.QueueDeclare(
			"hello", // name
			false,   // durable
			false,   // delete when unused
			false,   // exclusive
			false,   // no-wait
			nil,     // arguments
		)
		if err != nil {
			t.Error("Fail to declare channel:", err)
			return
		}

		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // arguments
		)
		if err != nil {
			t.Error("Fail to consume:", err)
			return
		}

		go func() {
			for m := range msgs {
				s := string(m.Body)
				if s == body {
					rcvbody <- s
				}
			}
			close(rcvdone)
		}()
	}

	// publisher
	{
		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		if err != nil {
			t.Error("Fail to dial:", err)
			return
		}

		defer func() {
			err = conn.Close()
			if err != nil {
				t.Log("Fail on close:", err)
			}
		}()

		ch, err := conn.Channel()
		if err != nil {
			t.Error("Fail to open channel:", err)
			return
		}

		q, err := ch.QueueDeclare(
			"hello", // queue
			false,   // durable
			false,   // delete when unused
			false,   // exclusive
			false,   // no-wait
			nil,     // arguments
		)
		if err != nil {
			t.Error("Fail to declare queue:", err)
			return
		}

		body = "Hello World"
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			},
		)
		if err != nil {
			t.Error("Fail to publish:", err)
			return
		}
	}

	timeout := 3 * time.Second
	select {
	case <-time.After(timeout):
		t.Errorf("No message after %s", timeout)
	case <-rcvbody:
		// found
	}
}
