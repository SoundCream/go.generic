package collection

func FromArray[T any](input []T) Collection[T] {
	return input
}

func From[T any](inputs ...T) Collection[T] {
	return inputs
}

func Map[T any, S any](c Collection[T], transform func(*T) S) []S {
	var result []S
	for i := range c {
		result = append(result, MapValue(&c[i], transform))
	}
	return result
}

func MapValue[T any, S any](input *T, transform func(*T) S) S {
	return transform(input)
}

func MapValueToPointer[T any, S any](input *T, transform func(*T) *S) *S {
	return transform(input)
}
