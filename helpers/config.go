package helpers

import (
	"reflect"
	_ "unsafe"

	_ "github.com/kelseyhightower/envconfig"
)

type configurationInfo struct {
	Name  string
	Alt   string
	Key   string
	Field reflect.Value
	Tags  reflect.StructTag
}

//go:linkname getConfigurationInfo github.com/kelseyhightower/envconfig.gatherInfo
func getConfigurationInfo(prefix string, config any) ([]configurationInfo, error)
