package linq

func (q Query[T]) Distinct() Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()
			set := make(map[interface{}]bool)

			return func() (item T, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if _, has := set[item]; !has {
						set[item] = true
						return
					}
				}

				return
			}
		},
	}
}
