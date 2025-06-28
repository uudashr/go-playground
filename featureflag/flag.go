package featureflag

var defaultFlagger = &Flagger{}

func SetDefaultFlagger(flagger *Flagger) {
	// TODO implement concurrency-safe logic
	defaultFlagger = flagger
}

type Flagger struct {
}

func (f *Flagger) Bool(name string, defaultValue bool) bool {
	// TODO implement feature flag logic
	return defaultValue
}

func (s *Flagger) String(name string, defaultValue string) string {
	// TODO implement feature flag logic
	return defaultValue
}

func Bool(name string, defaultValue bool) bool {
	// TODO implement concurrency-safe logic
	return defaultFlagger.Bool(name, defaultValue)
}

func String(name string, defaultValue string) string {
	// TODO implement concurrency-safe logic
	return defaultFlagger.String(name, defaultValue)
}

func StringWithOverride(name string, defaultValue string, override StringEvaluator) string {
	if override != nil {
		return override.String(name, defaultValue)
	}

	return String(name, defaultValue)
}

func BoolWithOverride(name string, defaultValue bool, override BoolEvaluator) bool {
	if override != nil {
		return override.Bool(name, defaultValue)
	}

	return Bool(name, defaultValue)
}

type StringEvaluator interface {
	String(name string, defaultValue string) string
}

type BoolEvaluator interface {
	Bool(name string, defaultValue bool) bool
}
