package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"reflect"
)

func main() {
	doDump()
}

func doMarshal() {
	v := User{
		Name:  "John Appleseed",
		Email: "john.appleseed@mail.com",
	}

	b, err := json.Marshal(v)
	if err != nil {
		log.Fatal("Marshal err:", err)
	}

	reconVal, err := unmarshal(b, User{})
	if err != nil {
		log.Fatal("unmarshal err:", err)
	}

	fmt.Printf("reconVal: %+v\n", reconVal)
}

func unmarshal(data []byte, v interface{}) (interface{}, error) {
	vv := v
	t := reflect.TypeOf(vv)
	refVal := reflect.New(t)
	fmt.Printf("Before: %v\n", reflect.TypeOf(vv))
	err := json.Unmarshal(data, refVal.Interface())
	if err != nil {
		return nil, err
	}

	fmt.Printf("After: %v\n", reflect.TypeOf(vv))
	return refVal.Elem().Interface(), nil
}

func doStructValue() {
	v := User{
		Name:  "John Appleseed",
		Email: "john.appleseed@mail.com",
	}

	t := reflect.TypeOf(v)
	refVal := reflect.New(t)
	fmt.Printf("Val: %+v\n", refVal.Elem().Interface())
}

func doSimpleValue() {
	v := "Hello"

	val := reflect.ValueOf(v)
	fmt.Printf("Value: %q\n", val.Interface())

	refVal := reflect.ValueOf(&v)
	refVal.Elem().Set(reflect.ValueOf("Hi"))
	fmt.Printf("Value: %q\n", v)

	refVal.Elem().SetString("Wow")
	fmt.Printf("Value: %q\n", v)
}

func doDump() {
	// v := LoginForm{
	// 	Email:    "john.appleseed@gmail.com",
	// 	Password: "secret",
	// 	valid:    true,
	// }

	// v := &LoginForm{
	// 	Email:    "john.appleseed@gmail.com",
	// 	Password: "secret",
	// 	valid:    true,
	// }

	// v := make(chan int, 10)

	// v := 10

	// var v io.Writer
	// v = &bytes.Buffer{}

	v := func(event EmailVerified) {

	}

	dump(v)
}

func dump(i interface{}) {
	t := reflect.TypeOf(i)

	fmt.Printf("- Type: %q\n", t)
	fmt.Printf("%2s- Name: %q\n", " ", t.Name())
	fmt.Printf("%2s- Kind: %q\n", " ", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Printf("%2s- Elem: %q\n", " ", t.Elem())
	case reflect.Struct:
		fmt.Printf("%2s- NumField: %d\n", " ", t.NumField())
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fmt.Printf("%4s- %s\n", " ", field.Name)
			fmt.Printf("%6s- Index: %v\n", " ", field.Index)
			fmt.Printf("%6s- Tag: %q\n", " ", field.Tag)
			fmt.Printf("%6s- PkgPath: %q\n", " ", field.PkgPath)
			fmt.Printf("%6s- Anonymous: %t\n", " ", field.Anonymous)
		}
	case reflect.Func:
		fmt.Printf("%2s- NumIn: %d\n", " ", t.NumIn())
		for i := 0; i < t.NumIn(); i++ {
			it := t.In(i)
			fmt.Printf("%4s- %q\n", " ", it)
		}

		fmt.Printf("%2s- NumOut: %d\n", " ", t.NumOut())
		for i := 0; i < t.NumOut(); i++ {
			it := t.Out(i)
			fmt.Printf("%4s- %q\n", " ", it)
		}
	}
}

type LoginForm struct {
	io.Writer
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
	valid    bool   `json:"valid"`
}

type User struct {
	Name  string
	Email string
}

type EmailVerified struct {
	UserID string
}
