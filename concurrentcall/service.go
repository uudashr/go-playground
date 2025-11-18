package concurrentcall

import (
	"log/slog"
	"sync"
)

type Service struct {
	callers []Caller
}

func NewService(callers []Caller) *Service {
	return &Service{
		callers: callers,
	}
}

func (svs *Service) Do() ([]string, error) {
	logger := slog.Default()

	var wg sync.WaitGroup
	wg.Add(len(svs.callers))
	outCh := make(chan string, len(svs.callers))

	for _, c := range svs.callers {
		go func() {
			defer wg.Done()

			out, err := c.Call()
			if err != nil {
				logger.Error("Failed to call", "error", err)
				return
			}

			outCh <- out
		}()
	}

	go func() {
		wg.Wait()
		close(outCh)
	}()

	var outputs []string
	for out := range outCh {
		outputs = append(outputs, out)
	}

	return outputs, nil
}

type Caller interface {
	Call() (string, error)
}
