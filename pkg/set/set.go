package set

import (
	"fmt"
	"strings"
)

type Set[T comparable] map[T]bool

func New[T comparable](initial []T) Set[T] {
	var set Set[T] = make(Set[T])
	for _, v := range initial {
		set[v] = true
	}
	return set
}

func (a Set[T]) String() string {
	builder := strings.Builder{}
	var first T

	for valueA := range a {
		first = valueA
		break
	}

	builder.WriteString(fmt.Sprintf("{ x : %T |", first))

	for valueA := range a {
		builder.WriteString(fmt.Sprintf(" %v ", valueA))
	}

	builder.WriteString("}")

	return builder.String()
}

func (a Set[T]) Sliceify() []T {
	var result []T = make([]T, 0)
	for valueA := range a {
		result = append(result, valueA)
	}
	return result
}

func (a Set[T]) Intersection(b Set[T]) Set[T] {
	var intersected Set[T] = make(Set[T])
	for valueB := range b {
		if a[valueB] {
			intersected[valueB] = true
		}
	}
	return intersected
}

func (a Set[T]) Union(b Set[T]) Set[T] {
	var unioned Set[T] = make(Set[T])

	for valueB := range b {
		unioned[valueB] = true
	}

	for valueA := range a {
		unioned[valueA] = true
	}

	return unioned
}

func (a Set[T]) Difference(b Set[T]) Set[T] {
	var intersected Set[T] = a.Intersection(b)
	var difference Set[T] = make(Set[T])

	for valueA := range a {
		if !intersected[valueA] {
			difference[valueA] = true
		}
	}

	return difference
}

func (a Set[T]) SymmetricDifference(b Set[T]) Set[T] {
	var intersected Set[T] = a.Intersection(b)
	var difference Set[T] = make(Set[T])

	for valueA := range a {
		if !intersected[valueA] {
			difference[valueA] = true
		}
	}

	for valueB := range b {
		if !intersected[valueB] {
			difference[valueB] = true
		}
	}

	return difference
}

func (a Set[T]) Subset(b Set[T]) bool {
	for valueA := range b {
		if !b[valueA] {
			return false
		}
	}
	return true
}

func (a Set[T]) Equal(b Set[T]) bool {
	return a.Subset(b) && b.Subset(a)
}
