package collection

type Collection[T any] []T

func (c Collection[T]) Map(transform func(*T) any) []any {
	var result []any
	for i := range c {
		result = append(result, MapValue(&c[i], transform))
	}
	return result
}

func (c Collection[T]) Filter(predicate func(*T) bool) Collection[T] {
	var result []T
	for i := range c {
		if predicate(&c[i]) {
			result = append(result, c[i])
		}
	}
	return result
}

func (c Collection[T]) Reduce(reduce func(current T, next T) T, initial T) T {
	result := initial
	for i := range c {
		result = reduce(result, c[i])
	}

	return result
}

func (c Collection[T]) Each(reduce func(value T)) {
	for i := range c {
		reduce(c[i])
	}
}

func (c Collection[T]) ToPtr(input []T) []*T {
	var result []*T
	for i := range input {
		result = append(result, ToPtr(input[i]))
	}
	return result
}

func ToPtr[T any](value T) *T {
	result := value
	return &result
}
