package featureflag

var Noop = &NoopFlagger{}

var defaultFlagger Flagger = Noop

func SetDefaultFlagger(flagger Flagger) {
	defaultFlagger = flagger
}

func DefaultFlagger() Flagger {
	if defaultFlagger == nil {
		return Noop
	}

	return defaultFlagger
}

type NoopFlagger struct {
}

func (nf *NoopFlagger) Bool(name string, defaultValue bool) bool {
	return defaultValue
}

func (nf *NoopFlagger) String(name string, defaultValue string) string {
	return defaultValue
}

func (nf *NoopFlagger) Int(name string, defaultValue int) int {
	return defaultValue
}

func Bool(name string, defaultValue bool) bool {
	return defaultFlagger.Bool(name, defaultValue)
}

func String(name string, defaultValue string) string {
	return defaultFlagger.String(name, defaultValue)
}

func Int(name string, defaultValue int) int {
	return defaultFlagger.Int(name, defaultValue)
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

func IntWithOverride(name string, defaultValue int, override IntEvaluator) int {
	if override != nil {
		return override.Int(name, defaultValue)
	}

	return Int(name, defaultValue)
}

type Flagger interface {
	StringEvaluator
	BoolEvaluator
	IntEvaluator
}

type StringEvaluator interface {
	String(name string, defaultValue string) string
}

type BoolEvaluator interface {
	Bool(name string, defaultValue bool) bool
}

type IntEvaluator interface {
	Int(name string, defaultValue int) int
}
