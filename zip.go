package linq

func Zip[T1, T2, T3 any](q Query[T1], q2 Query[T2],
	resultSelector func(T1, T2) T3) Query[T3] {

	return Query[T3]{
		Iterate: func() Iterator[T3] {
			next1 := q.Iterate()
			next2 := q2.Iterate()

			return func() (item T3, ok bool) {
				item1, ok1 := next1()
				item2, ok2 := next2()

				if ok1 && ok2 {
					return resultSelector(item1, item2), true
				}

				return item, false
			}
		},
	}
}
