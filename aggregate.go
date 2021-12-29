package linq

func (q Query[T]) Aggregate(f func(T, T) T) T {
	next := q.Iterate()

	result, any := next()
	if !any {
		var t T
		return t
	}

	for current, ok := next(); ok; current, ok = next() {
		result = f(result, current)
	}

	return result
}

func (q Query[T]) AggregateWithSeed(seed T, f func(T, T) T) T {

	next := q.Iterate()
	result := seed

	for current, ok := next(); ok; current, ok = next() {
		result = f(result, current)
	}

	return result
}

func (q Query[T]) AggregateWithSeedBy(seed T,
	f func(T, T) T,
	resultSelector func(T) T) T {

	next := q.Iterate()
	result := seed

	for current, ok := next(); ok; current, ok = next() {
		result = f(result, current)
	}

	return resultSelector(result)
}
