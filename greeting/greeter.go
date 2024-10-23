package greeting

type Greeter interface {
	Greet(name string) string
}

func Greet(g Greeter, name string) string {
	// return "Hello, " + name + "!"
	return g.Greet(name)
}
