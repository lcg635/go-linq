package linq

type Group[K comparable, T any] struct {
	Key   K
	Group []T
}

func GroupBy[K comparable, T1, T2 any](q Query[T1], keySelector func(T1) K,
	elementSelector func(T1) T2) Query[Group[K, T2]] {
	return Query[Group[K, T2]]{
		func() Iterator[Group[K, T2]] {
			next := q.Iterate()
			set := make(map[K][]T2)

			for item, ok := next(); ok; item, ok = next() {
				key := keySelector(item)
				set[key] = append(set[key], elementSelector(item))
			}

			len := len(set)
			idx := 0
			groups := make([]Group[K, T2], len)
			for k, v := range set {
				groups[idx] = Group[K, T2]{k, v}
				idx++
			}

			index := 0

			return func() (item Group[K, T2], ok bool) {
				ok = index < len
				if ok {
					item = groups[index]
					index++
				}

				return
			}
		},
	}
}
