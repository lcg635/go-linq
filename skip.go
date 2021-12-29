package linq

func (q Query[T]) Skip(count int) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()
			n := count

			return func() (item T, ok bool) {
				for ; n > 0; n-- {
					item, ok = next()
					if !ok {
						return
					}
				}

				return next()
			}
		},
	}
}

func (q Query[T]) SkipWhile(predicate func(T) bool) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()
			ready := false

			return func() (item T, ok bool) {
				for !ready {
					item, ok = next()
					if !ok {
						return
					}

					ready = !predicate(item)
					if ready {
						return
					}
				}

				return next()
			}
		},
	}
}

func (q Query[T]) SkipWhileIndexed(predicate func(int, T) bool) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()
			ready := false
			index := 0

			return func() (item T, ok bool) {
				for !ready {
					item, ok = next()
					if !ok {
						return
					}

					ready = !predicate(index, item)
					if ready {
						return
					}

					index++
				}

				return next()
			}
		},
	}
}
