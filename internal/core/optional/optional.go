package optional

type Optional[T any] struct {
	value     T
	isPresent bool
}

func NewOptional[T any](value T) Optional[T] {
	return Optional[T]{
		value:     value,
		isPresent: true,
	}
}

func NewEmptyOptional[T any]() Optional[T] {
	return Optional[T]{
		isPresent: false,
	}
}

func NewOptionalFromPointer[T any](value *T) Optional[T] {
	if value == nil {
		return NewEmptyOptional[T]()
	}
	return NewOptional(*value)
}

func (o Optional[T]) Value() T {
	if o.isPresent {
		return o.value
	}

	var zero T
	return zero
}

func (o Optional[T]) IsPresent() bool {
	return o.isPresent
}

func (o Optional[T]) ToPointer() *T {
	if o.isPresent {
		return &o.value
	}
	return nil
}
