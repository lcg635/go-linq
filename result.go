package linq

func (q Query[T]) All(predicate func(T) bool) bool {
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		if !predicate(item) {
			return false
		}
	}

	return true
}

func (q Query[T]) Any() bool {
	_, ok := q.Iterate()()
	return ok
}

func (q Query[T]) AnyWith(predicate func(T) bool) bool {
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		if predicate(item) {
			return true
		}
	}

	return false
}

func (q Query[T]) Count() (r int) {
	next := q.Iterate()

	for _, ok := next(); ok; _, ok = next() {
		r++
	}

	return
}

func (q Query[T]) CountWith(predicate func(T) bool) (r int) {
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		if predicate(item) {
			r++
		}
	}

	return
}

func (q Query[T]) First() interface{} {
	item, _ := q.Iterate()()
	return item
}

func (q Query[T]) FirstWith(predicate func(T) bool) (t T) {
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		if predicate(item) {
			return item
		}
	}

	return
}

func (q Query[T]) ForEach(action func(T)) {
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		action(item)
	}
}

func (q Query[T]) ForEachIndexed(action func(int, T)) {
	next := q.Iterate()
	index := 0

	for item, ok := next(); ok; item, ok = next() {
		action(index, item)
		index++
	}
}

func (q Query[T]) Last() (r T) {
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		r = item
	}

	return
}

func (q Query[T]) LastWith(predicate func(T) bool) (r T) {
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		if predicate(item) {
			r = item
		}
	}

	return
}

func (q Query[T]) Results() []T {
	next := q.Iterate()

	var results []T
	for item, ok := next(); ok; item, ok = next() {
		results = append(results, item)
	}
	return results
}

func Contains[T comparable](q Query[T], value T) bool {
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		if item == value {
			return true
		}
	}

	return false
}

func ToMap[K comparable, V any](q Query[KV[K, V]]) map[K]V {
	m := make(map[K]V)
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		m[item.Key] = item.Value
	}
	return m
}

func ToGroupMap[K comparable, V any](q Query[Group[K, V]]) map[K][]V {
	m := make(map[K][]V)
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		m[item.Key] = item.Group
	}
	return m
}
