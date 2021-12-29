package linq

func (q Query[T]) Union(q2 Query[T]) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()
			next2 := q2.Iterate()

			set := make(map[interface{}]bool)
			use1 := true

			return func() (item T, ok bool) {
				if use1 {
					for item, ok = next(); ok; item, ok = next() {
						if _, has := set[item]; !has {
							set[item] = true
							return
						}
					}

					use1 = false
				}

				for item, ok = next2(); ok; item, ok = next2() {
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
