package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)

			// buf := make([]byte, 1024)
			// n := runtime.Stack(buf, false)
			// fmt.Printf("Stack trace:\n%s\n", buf[:n])

			// captureStackTrace()

			fmt.Printf("Stack trace:\n%s\n", string(debug.Stack()))
		}
	}()
	// fmt.Println("Let's close")
	// close()
	// fmt.Println("Closed")

	var p *MyStruct
	p.DoSomething() // => panic if DoSomething() accesses p.Value
}

func captureStackTrace() {
	buf := make([]byte, 1024)
	n := runtime.Stack(buf, false)
	fmt.Printf("Stack trace:\n%s\n", buf[:n])
}

func close() {
	var src *Resource

	src.Close()
}

type Resource struct {
}

func (r *Resource) Close() {
	fmt.Println("Closing resource")
	panic("oops")
}

type MyStruct struct {
	Value int
}

func (m *MyStruct) DoSomething() {
	// Using m.Value on a nil receiver => panic
	fmt.Println(m.Value)
}
