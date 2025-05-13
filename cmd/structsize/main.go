package main

import (
	"fmt"
	"unsafe"
)

type Inner struct {
	A int64
	B byte
}

type Outer struct {
	X int
	Y Inner
	Z float64
}

func main() {
	fmt.Println("Size of Inner:", unsafe.Sizeof(Inner{})) // Will include padding
	fmt.Println("Size of Outer:", unsafe.Sizeof(Outer{})) // Includes size of Inner
}
