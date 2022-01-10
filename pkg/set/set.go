package set

import (
	"fmt"
	"strings"
)

// Type Set[T] is a generic mathematical set data type.
type Set[T comparable] map[T]bool

// New creates a new set from an initial slice
func New[T comparable](initial []T) Set[T] {
	var set Set[T] = make(Set[T])
	for _, v := range initial {
		set[v] = true
	}
	return set
}

// String returns a pretty looking string representation of the set.
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

// Intersection performs set intersection on two sets.
func (a Set[T]) Intersection(b Set[T]) Set[T] {
	var intersected Set[T] = make(Set[T])
	for valueB := range b {
		if a[valueB] {
			intersected[valueB] = true
		}
	}
	return intersected
}

// Union performs set union on two sets.
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

// Difference performs set difference on two sets.
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

// SymmetricDifference performs symmetric set difference on two sets.
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

// Subset returns true if a is a subset of b, and false otherwise.
func (a Set[T]) Subset(b Set[T]) bool {
	for valueA := range a {
		if !b[valueA] {
			return false
		}
	}
	return true
}

// Equal returns true if two sets are equal, false otherwise.
func (a Set[T]) Equal(b Set[T]) bool {
	return a.Subset(b) && b.Subset(a)
}
