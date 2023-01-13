package mockplay

//go:generate mockery --name=Greeter --log-level=warn
type Greeter interface {
	Greet(msg string) string
}

func Hello(g Greeter) string {
	return g.Greet("Hello")
}
