package linq

type Iterator[T any] func() (item T, ok bool)

type Query[T any] struct {
	Iterate func() Iterator[T]
}

type Iterable[T any] interface {
	Iterate() Iterator[T]
}

type KV[K, V any] struct {
	Key   K
	Value V
}

func FromArray[T any](source []T) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			index := 0
			len := len(source)

			return func() (item T, ok bool) {
				ok = index < len
				if ok {
					item = source[index]
					index++
				}
				return
			}
		},
	}
}

func FromString(source string) Query[rune] {
	runes := []rune(source)
	len := len(runes)

	return Query[rune]{
		Iterate: func() Iterator[rune] {
			index := 0

			return func() (item rune, ok bool) {
				ok = index < len
				if ok {
					item = runes[index]
					index++
				}

				return
			}
		},
	}
}

func FromMap[K comparable, V any](source map[K]V) Query[KV[K, V]] {
	len := len(source)

	return Query[KV[K, V]]{
		Iterate: func() Iterator[KV[K, V]] {
			index := 0
			keys := make([]K, 0, len)
			for k := range source {
				keys = append(keys, k)
			}

			return func() (item KV[K, V], ok bool) {
				ok = index < len
				if ok {
					key := keys[index]
					item = KV[K, V]{Key: key, Value: source[key]}
					index++
				}
				return
			}
		},
	}
}

func FromIterable[T any](source Iterable[T]) Query[T] {
	return Query[T]{
		Iterate: source.Iterate,
	}
}

func Range[T int | int32 | int64](start, count T) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			var index T
			current := start

			return func() (item T, ok bool) {
				if index >= count {
					return 0, false
				}

				item, ok = current, true

				index++
				current++
				return
			}
		},
	}
}

func Repeat[T any](value T, count int) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			index := 0

			return func() (item T, ok bool) {
				if index >= count {
					var t T
					return t, false
				}
				item, ok = value, true
				index++
				return
			}
		},
	}
}
