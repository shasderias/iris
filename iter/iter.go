package iter

type Iterator[T any] interface {
	Iterate() Iter[T]
}

type Iter[T any] interface {
	Next() (T, bool)
}

func Map[T, U any](iter Iter[T], mapFn func(T) U) []U {
	out := make([]U, 0)
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		out = append(out, mapFn(v))
	}
	return out
}

func ForEach[T any](iter Iter[T], fn func(T)) {
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		fn(v)
	}
}

func MapSlice[T, U any](slice []T, mapFn func(T) U) []U {
	out := make([]U, len(slice))
	for i, v := range slice {
		out[i] = mapFn(v)
	}
	return out
}

func SliceHas[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}