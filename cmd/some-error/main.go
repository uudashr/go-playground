package main

import (
	"errors"
	"fmt"
)

func main() {
	var err error
	err = errors.Join(err, nil)
	err = errors.Join(err, errors.New("first error"))
	err = errors.Join(err, errors.New("second error"))
	err = errors.Join(err, nil)
	err = errors.Join(err, errors.New("third error"))
	if err != nil {
		fmt.Println("Got error:", err)
	} else {
		fmt.Println("All good")
	}
}
