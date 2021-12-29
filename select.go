package linq

func (q Query[T]) Select(selector func(T) T) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()

			return func() (item T, ok bool) {
				var it T
				it, ok = next()
				if ok {
					item = selector(it)
				}
				return
			}
		},
	}
}

func (q Query[T]) SelectIndexed(selector func(int, T) T) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()
			index := 0

			return func() (item T, ok bool) {
				var it T
				it, ok = next()
				if ok {
					item = selector(index, it)
					index++
				}

				return
			}
		},
	}
}
