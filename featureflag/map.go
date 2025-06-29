package featureflag

type MapFlagger struct {
	flags map[string]any
}

func NewMapFlagger(flags map[string]any) *MapFlagger {
	return &MapFlagger{flags: flags}
}

func (m *MapFlagger) Bool(name string, defaultValue bool) bool {
	return mapValue(m.flags, name, defaultValue)
}

func (m *MapFlagger) String(name string, defaultValue string) string {
	return mapValue(m.flags, name, defaultValue)
}

func (m *MapFlagger) Int(name string, defaultValue int) int {
	return mapValue(m.flags, name, defaultValue)
}

func mapValue[T any](m map[string]any, name string, defaultValue T) T {
	if val, ok := m[name]; ok {
		return val.(T)
	}

	return defaultValue
}
