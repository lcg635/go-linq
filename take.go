package linq

func (q Query[T]) Take(count int) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()
			n := count

			return func() (item T, ok bool) {
				if n <= 0 {
					return
				}

				n--
				return next()
			}
		},
	}
}

func (q Query[T]) TakeWhile(predicate func(T) bool) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()
			done := false

			return func() (item T, ok bool) {
				if done {
					return
				}

				item, ok = next()
				if !ok {
					done = true
					return
				}

				if predicate(item) {
					return
				}

				var t T
				done = true
				return t, false
			}
		},
	}
}

func (q Query[T]) TakeWhileIndexed(predicate func(int, T) bool) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()
			done := false
			index := 0

			return func() (item T, ok bool) {
				if done {
					return
				}

				item, ok = next()
				if !ok {
					done = true
					return
				}

				if predicate(index, item) {
					index++
					return
				}

				var t T
				done = true
				return t, false
			}
		},
	}
}
