package linq

func (q Query[T]) Reverse() Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()

			items := []T{}
			for item, ok := next(); ok; item, ok = next() {
				items = append(items, item)
			}

			index := len(items) - 1
			return func() (item T, ok bool) {
				if index < 0 {
					return
				}

				item, ok = items[index], true
				index--
				return
			}
		},
	}
}
