package main

import (
	"fmt"
	"reflect"
)

func TypeOf[T any]() reflect.Type {
	var t T
	return reflect.TypeOf(t)
}

func main() {
	t := TypeOf[OrderCompleted]()
	fmt.Println(t)

	fn := func(event OrderCompleted) {
		fmt.Println("OrderCompleted:", event)
	}

	h, err := newFnHandler(fn)
	if err != nil {
		panic(err)
	}

	h.Handle(OrderCompleted{OrderID: "123"})
}

type OrderCompleted struct {
	OrderID string
}

type CustomerSuspended struct {
	CustomerID string
}

type Handle interface {
	Handle(event interface{})
}

type HandlerFunc func(event interface{})

func (f HandlerFunc) Handle(event interface{}) {
	f(event)
}

type fnHandler struct {
	fn      interface{}
	argType reflect.Type
}

func newFnHandler(fn interface{}) (*fnHandler, error) {
	if fn == nil {
		return nil, fmt.Errorf("fn is nil")
	}

	fnType := reflect.TypeOf(fn)
	if fnType.Kind() != reflect.Func {
		return nil, fmt.Errorf("fn is not a function")
	}

	if fnType.NumIn() != 1 {
		return nil, fmt.Errorf("fn must have exactly one argument")
	}

	if fnType.NumOut() != 0 {
		return nil, fmt.Errorf("fn must have no return values")
	}

	argType := fnType.In(0)
	if argType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("fn input parameter should be a struct (got: %v)", fnType.In(0).Kind())
	}

	return &fnHandler{
		fn:      fn,
		argType: argType,
	}, nil
}

func (h *fnHandler) Handle(event interface{}) {
	invokeHandler(h.fn, event)
}

func invokeHandler(fn interface{}, event interface{}) {
	handler := reflect.ValueOf(fn)
	handler.Call([]reflect.Value{reflect.ValueOf(event)})
}
