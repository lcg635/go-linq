package linq

import "testing"

type Foo interface {
	CC()
}

type X struct {
	A int
	B string
	C bool
}

func TestArray(t *testing.T) {
	x := FromArray([]X{
		{A: 1, B: "a", C: false},
		{A: 2, B: "b", C: false},
		{A: 3, B: "a", C: true},
		{A: 4, B: "d", C: false},
		{A: 5, B: "b", C: true},
	})
	g := GroupBy(x,
		func(t X) bool {
			return t.C
		}, func(t X) int {
			return t.A
		})
	t.Log(ToGroupMap(g))
}

func TestAggregate(t *testing.T) {
	x := FromArray([]int{1, 2, 3, 4, 5}).
		Aggregate(func(t1, t2 int) int {
			return t1 + t2
		})
	t.Log(x)
	y := FromArray([]int{1, 2, 3, 4, 4, 4, 4}).Distinct()
	t.Log(y.Results())
}

func TestMap(t *testing.T) {
	x := FromMap(map[string]int{"a": 1, "b": 2, "c": 1, "d": 2}).
		Where(func(t KV[string, int]) bool {
			return true
		})
	t.Log(ToMap(x))
	it := GroupBy(x, func(t KV[string, int]) int {
		return t.Value
	}, func(t KV[string, int]) string {
		return t.Key
	})
	t.Log(ToGroupMap(it))
}

func TestRepeat(t *testing.T) {
	x := Repeat(2, 10).Results()
	t.Log(x)
}

func TestRange(t *testing.T) {
	x := Range(2, 10).Results()
	t.Log(x)
}
