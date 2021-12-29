package linq

func (q Query[T]) Where(predicate func(T) bool) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()

			return func() (item T, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if predicate(item) {
						return
					}
				}
				return
			}
		},
	}
}

func (q Query[T]) WhereIndexed(predicate func(int, T) bool) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()
			index := 0

			return func() (item T, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if predicate(index, item) {
						index++
						return
					}

					index++
				}

				return
			}
		},
	}
}
