package property

// String defines the contract for a string property.
type String interface {
	// Set updates the property value and fires an event, if the new value is different than the old value.
	Set(v string)

	// Get returns the current property value
	Get() string

	// Bind connects the given pointer to a string with the value. The first time, the value from dst is read and
	// populates the property. Afterwards the direction is always the opposite, and updates to the property
	// will update the dst.
	Bind(dst *string)

	// Observe registers a callback which is fired, if the value has been set. It is not fire, if the value has not
	// been changed, e.g. if setting the same string.
	Observe(onDidSet func(old, new string)) Func
}

// NewString creates a new self-contained property.
func NewString() String {
	return &stringProperty{absProperty: newAbsProperty()}
}

type stringProperty struct {
	*absProperty
}

func (s *stringProperty) Set(v string) {
	s.absProperty.Set(v)
}

func (s *stringProperty) Get() string {
	if s.absProperty.Get() == nil {
		return ""
	}

	return s.absProperty.Get().(string)
}

// TODO when and where to unbind?
func (s *stringProperty) Bind(dst *string) {
	s.absProperty.Observe(func(old, new interface{}) {
		*dst = new.(string)
	})
	// TODO unclear if this is a good idea, it is the only time, it will make this way
	s.Set(*dst)
}

func (s *stringProperty) Observe(onDidSet func(old, new string)) Func {
	return s.absProperty.Observe(func(old, new interface{}) {
		if old == nil {
			old = ""
		}

		if old != new {
			onDidSet(old.(string), new.(string))
		}
	})
}
