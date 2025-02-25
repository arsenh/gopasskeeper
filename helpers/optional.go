package helpers

import "gopasskeeper/constants"

type Optional[T any] struct {
	value    T
	hasValue bool
}

func (o Optional[T]) HasValue() bool {
	return o.hasValue
}

func (o Optional[T]) Get() (T, bool) {
	return o.value, o.hasValue
}

func (o Optional[T]) MustGet() T {
	if !o.hasValue {
		panic(constants.ErrAccessOptionalParam)
	}
	return o.value
}

func (o *Optional[T]) Reset() {
	var zero T
	o.value = zero
	o.hasValue = false
}

func NewOptional[T comparable](value T) Optional[T] {
	var zero T
	return Optional[T]{
		value:    value,
		hasValue: value != zero, // Check if value is not the zero value
	}
}
