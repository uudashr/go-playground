package main

import "fmt"

func main() {
	reduceSliceOut()

	// Questions:
	// What if the capacity is higher than the length
}

func reduceSlice() {
	// Declaring a slice
	numbers := []int{10, 20, 30, 40, 50, 90, 60}
	fmt.Println("Original slice:", numbers, len(numbers), cap(numbers))
	fmt.Printf("Original slice address: %p\n", &numbers)

	var index int = 3

	// Get the element at the provided index in the slice
	elem := numbers[index]

	// Using append function to combine two slices
	// first slice is the slice of all the elements before the given index
	// second slice is the slice of all the elements after the given index
	// append function appends the second slice to the end of the first slice
	// returning a slice, so we store it in the form of a slice
	numbers = append(numbers[:index], numbers[index+1:]...)
	out := numbers

	out[0] = 100

	fmt.Println("Original slice after deleting elements:", numbers, len(numbers), cap(numbers))
	fmt.Printf("Original slice address: %p\n", &numbers)
	fmt.Printf("The element %d at index %d was deleted.\n", elem, index)
	fmt.Println("Slice output:", out, len(out), cap(out))
	fmt.Printf("Slice output address: %p\n", &out)

	// Observation:
	// - Ok if the original slice is defined inside the function, the slice is also the output slice
	// - Might not ok, if we found that slice is not passed to function as argment. Because there are possibilities for the caller expect the argument for not mutated
	// - The address of the append result is different from the original if we use declaration assigment
	// - Any attempt to slice modification will affect the slice source of the assignment
}

func reduceSliceOut() {
	// Declaring a slice
	numbers := []int{10, 20, 30, 40, 50, 90, 60}
	fmt.Println("Original slice:", numbers, len(numbers), cap(numbers))
	fmt.Printf("Original slice address: %p\n", &numbers)

	var index int = 3

	// Get the element at the provided index in the slice
	elem := numbers[index]

	// Using append function to combine two slices
	// first slice is the slice of all the elements before the given index
	// second slice is the slice of all the elements after the given index
	// append function appends the second slice to the end of the first slice
	// returning a slice, so we store it in the form of a slice
	out := append(numbers[:index], numbers[index+1:]...)

	out[0] = 100

	fmt.Println("Original slice after deleting elements:", numbers, len(numbers), cap(numbers))
	fmt.Printf("Original slice address: %p\n", &numbers)
	fmt.Printf("The element %d at index %d was deleted.\n", elem, index)
	fmt.Println("Slice output:", out, len(out), cap(out))
	fmt.Printf("Slice output address: %p\n", &out)

	// Observation:
	// - Original slice elemets is mutated, but the length is not changed
	// - It might give false perception if we use using the out and leave the original slice,
	//    include if the slice is taken from function arg because they will be mutated
}

func addSlice() {
	// Declaring a slice
	numbers := []int{10, 20, 30, 40, 50, 90, 60}
	fmt.Println("Original slice:", numbers, len(numbers), cap(numbers))
	fmt.Printf("Original slice address: %p\n", &numbers)

	numbers = append(numbers, 200)

	out := numbers

	out[0] = 100

	fmt.Println("Original slice after adding elements:", numbers, len(numbers), cap(numbers))
	fmt.Printf("Original slice address: %p\n", &numbers)
	fmt.Println("Slice output:", out, len(out), cap(out))
	fmt.Printf("Slice output address: %p\n", &out)

	// Observation::
	// - The length is increase and the capacity increase from 7 to 14 (2x)
	// - It still use the same address
}

func initSlice() {
	var numbers1 []int
	numbers2 := []int{}
	numbers3 := make([]int, 0)
	numbers4 := make([]int, 0, 10)
	numbers5 := make([]int, 10)
	numbers6 := make([]int, 10, 20)

	fmt.Println("Slice 1:", numbers1, len(numbers1), cap(numbers1))
	fmt.Println("Slice 2:", numbers2, len(numbers2), cap(numbers2))
	fmt.Println("Slice 3:", numbers3, len(numbers3), cap(numbers3))
	fmt.Println("Slice 4:", numbers4, len(numbers4), cap(numbers4))
	fmt.Println("Slice 5:", numbers5, len(numbers5), cap(numbers5))
	fmt.Println("Slice 6:", numbers6, len(numbers6), cap(numbers6))

	// Observation:
	// - It prints value in its length

}
