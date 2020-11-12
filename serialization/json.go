package serialization

import (
	"encoding/json"
	"errors"
	"reflect"
)

type Deserializer struct {
	valSamples map[string]interface{}
}

func (d *Deserializer) Register(v interface{}) {
	name := typeName(v)
	d.RegisterName(name, v)
}

func (d *Deserializer) RegisterName(name string, v interface{}) {
	if d.valSamples == nil {
		d.valSamples = make(map[string]interface{})
	}

	d.valSamples[name] = v
}

func (d *Deserializer) Deserialize(name string, data []byte) (interface{}, error) {
	if d.valSamples == nil {
		return nil, errors.New("unrecognized event")
	}

	v := d.valSamples[name]
	refTyp := reflect.TypeOf(v)
	refPtrVal := reflect.New(refTyp)

	err := json.Unmarshal(data, refPtrVal.Interface())
	if err != nil {
		return nil, err
	}

	return refPtrVal.Elem().Interface(), nil
}

func typeName(i interface{}) string {
	return reflect.TypeOf(i).Name()
}
