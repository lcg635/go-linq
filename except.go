package linq

func (q Query[T]) Except(q2 Query[T]) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()

			next2 := q2.Iterate()
			set := make(map[interface{}]bool)
			for i, ok := next2(); ok; i, ok = next2() {
				set[i] = true
			}

			return func() (item T, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if _, has := set[item]; !has {
						return
					}
				}

				return
			}
		},
	}
}

func (q Query[T]) ExceptBy(
	q2 Query[T],
	selector func(T) T,
) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()

			next2 := q2.Iterate()
			set := make(map[interface{}]bool)
			for i, ok := next2(); ok; i, ok = next2() {
				s := selector(i)
				set[s] = true
			}

			return func() (item T, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					s := selector(item)
					if _, has := set[s]; !has {
						return
					}
				}

				return
			}
		},
	}
}
