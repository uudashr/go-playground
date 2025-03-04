package diginjection_test

import (
	"encoding/json"
	"log"
	"os"

	"go.uber.org/dig"
)

type Config struct {
	Prefix string
}

func ExampleContainer() {
	c := dig.New()

	err := c.Provide(func() (*Config, error) {
		var cfg Config
		err := json.Unmarshal([]byte(`{ "prefix": "[jungle] " }`), &cfg)
		if err != nil {
			return nil, err
		}

		return &cfg, nil
	})
	if err != nil {
		panic(err)
	}

	err = c.Provide(func(cfg *Config) *log.Logger {
		return log.New(os.Stdout, cfg.Prefix, 0)
	})
	if err != nil {
		panic(err)
	}

	err = c.Invoke(func(l *log.Logger) {
		l.Print("You've been invoked")
	})
	if err != nil {
		panic(err)
	}

	// Output: [jungle] You've been invoked
}
