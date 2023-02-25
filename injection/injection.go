package injection

import (
	"errors"
	"fmt"
	"reflect"
)

type Label string

type Registry struct {
	valueProviders map[Label]valueProviders
}

func NewRegistry() *Registry {
	return &Registry{
		valueProviders: make(map[Label]valueProviders),
	}
}

func (r *Registry) Provide(v any, label Label, argLabels ...Label) error {
	valType := reflect.TypeOf(v)
	if valType.Kind() == reflect.Func {
		return r.ProvideFunc(v, label, argLabels...)
	}

	r.ProvideVal(v, label)
	return nil
}

func (r *Registry) ProvideVal(v any, label Label) {
	r.valueProviders[label] = &staticValue{val: v}
}

func (r *Registry) ProvideFunc(fn any, label Label, argLabels ...Label) error {
	rt := reflect.TypeOf(fn)
	if rt.Kind() != reflect.Func {
		return errors.New("ProvideFunc expecting fn argument ")
	}

	if rt.NumIn() != len(argLabels) {
		return fmt.Errorf("ProvideFunc expecting %d labels, got %d", rt.NumIn(), len(argLabels))
	}

	r.valueProviders[label] = &funcValue{
		fn:        fn,
		argLabels: argLabels,
	}

	return nil
}

func (r *Registry) Resolve(label Label) (any, error) {
	p := r.valueProviders[label]
	if p == nil {
		return nil, fmt.Errorf("no value provided for label %s", label)
	}

	return p.provideValue(r)
}

func (r *Registry) InjectFunc(fn any, labels ...Label) error {
	fnType := reflect.TypeOf(fn)
	if fnType.Kind() != reflect.Func {
		return errors.New("expecting fn argument ")
	}

	if fnType.NumIn() != len(labels) {
		return fmt.Errorf("expecting %d labels, got %d", fnType.NumIn(), len(labels))
	}

	fnVal := reflect.ValueOf(fn)
	argValues := make([]reflect.Value, len(labels))
	for i := 0; i < fnType.NumIn(); i++ {
		v, err := r.Resolve(labels[i])
		if err != nil {
			return err
		}

		argValues[i] = reflect.ValueOf(v)
	}
	fnVal.Call(argValues)

	return nil
}

func (r *Registry) Inject(v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer {
		return errors.New("expecting pointer argument ")
	}

	elem := rv.Elem()
	if elem.Kind() != reflect.Struct {
		return errors.New("expecting pointer of struct argument")
	}

	elemType := elem.Type()
	for i := 0; i < elemType.NumField(); i++ {
		ft := elemType.Field(i)
		tagVal := ft.Tag.Get("injection")
		val, err := r.Resolve(Label(tagVal))
		if err != nil {
			return err
		}

		f := elem.Field(i)
		f.Set(reflect.ValueOf(val))
	}

	return nil
}

type Resolver interface {
	Resolve(label Label) (any, error)
}

type valueProviders interface {
	provideValue(r Resolver) (any, error)
}

type staticValue struct {
	val any
}

func (sv *staticValue) provideValue(r Resolver) (any, error) {
	return sv.val, nil
}

type funcValue struct {
	fn        any
	argLabels []Label
}

func (gv *funcValue) provideValue(r Resolver) (any, error) {
	rv := reflect.ValueOf(gv.fn)
	if rv.Kind() != reflect.Func {
		panic("funcValue.fn expecting fn argument")
	}

	argVals := make([]reflect.Value, len(gv.argLabels))
	for i, argLabel := range gv.argLabels {
		// TODO: we need to detect cyclic dependencies
		v, err := r.Resolve(argLabel)
		if err != nil {
			return nil, err
		}

		argVals[i] = reflect.ValueOf(v)
	}
	return rv.Call(argVals)[0].Interface(), nil

}

type Valuer[T any] interface {
	Value() (T, error)
}
