package linq

func (q Query[T]) Append(item T) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()
			appended := false

			return func() (T, bool) {
				i, ok := next()
				if ok {
					return i, ok
				}

				if !appended {
					appended = true
					return item, true
				}

				var t T
				return t, false
			}
		},
	}
}

func (q Query[T]) Concat(q2 Query[T]) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()
			next2 := q2.Iterate()
			use1 := true

			return func() (item T, ok bool) {
				if use1 {
					item, ok = next()
					if ok {
						return
					}

					use1 = false
				}

				return next2()
			}
		},
	}
}

func (q Query[T]) Prepend(item T) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()
			prepended := false

			return func() (T, bool) {
				if prepended {
					return next()
				}

				prepended = true
				return item, true
			}
		},
	}
}
