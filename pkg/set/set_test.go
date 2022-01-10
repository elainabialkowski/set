package set_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/ElainaBialkowski/set/pkg/set"
)

var BenchmarkInputs = map[int]struct {
	a set.Set[int]
	b set.Set[int]
}{
	100:    {set.New(RandomIntegers(100)), set.New(RandomIntegers(100))},
	1000:   {set.New(RandomIntegers(1000)), set.New(RandomIntegers(1000))},
	10000:  {set.New(RandomIntegers(10000)), set.New(RandomIntegers(10000))},
	100000: {set.New(RandomIntegers(100000)), set.New(RandomIntegers(100000))},
}

func Benchmark_Intersection(t *testing.B) {
	for k, v := range BenchmarkInputs {
		t.Run(fmt.Sprintf("input_size_%d", k), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				v.a.Intersection(v.b)
			}
		})
	}
}

func Benchmark_Union(t *testing.B) {
	for k, v := range BenchmarkInputs {
		t.Run(fmt.Sprintf("input_size_%d", k), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				v.a.Union(v.b)
			}
		})
	}
}

func Benchmark_Difference(t *testing.B) {
	for k, v := range BenchmarkInputs {
		t.Run(fmt.Sprintf("input_size_%d", k), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				v.a.Difference(v.b)
			}
		})
	}
}

func Test_Intersection(t *testing.T) {
	var a set.Set[int] = set.New([]int{1, 9, 6, 4, 5})
	var b set.Set[int] = set.New([]int{1, 4, 5})
	var expected set.Set[int] = set.New([]int{1, 4, 5})

	if !a.Intersection(b).Equal(expected) {
		t.Fail()
	}

}

func Test_Union(t *testing.T) {
	var a set.Set[int] = set.New([]int{1, 9, 6, 4, 5})
	var b set.Set[int] = set.New([]int{1, 4, 5})
	var expected set.Set[int] = set.New([]int{1, 9, 6, 4, 5})

	if !a.Union(b).Equal(expected) {
		t.Fail()
	}
}

func Test_Difference(t *testing.T) {
	var a set.Set[int] = set.New([]int{1, 9, 6, 4, 5})
	var b set.Set[int] = set.New([]int{1, 4, 5})
	var expected set.Set[int] = set.New([]int{9, 6})

	if !a.Difference(b).Equal(expected) {
		t.Fail()
	}
}

func Test_Symmetric_Difference(t *testing.T) {
	var a set.Set[int] = set.New([]int{1, 9, 6, 4, 5})
	var b set.Set[int] = set.New([]int{1, 4, 5, 17, 15})
	var expected set.Set[int] = set.New([]int{9, 6, 17, 15})

	if !a.SymmetricDifference(b).Equal(expected) {
		t.Fail()
	}
}

func Test_Subset_True(t *testing.T) {
	var a set.Set[int] = set.New([]int{1, 9, 6, 4, 5})
	var b set.Set[int] = set.New([]int{1, 4, 5})

	if !b.Subset(a) {
		t.Fail()
	}
}

func Test_Subset_False(t *testing.T) {
	var a set.Set[int] = set.New([]int{1, 9, 6, 4, 5})
	var b set.Set[int] = set.New([]int{1, 4, 5, 2})

	if b.Subset(a) {
		t.Fail()
	}
}

func Test_Equal(t *testing.T) {
	var a set.Set[int] = set.New([]int{1, 9, 6, 4, 5})
	var b set.Set[int] = set.New([]int{1, 4, 5, 9, 6})

	if !b.Equal(a) {
		t.Fail()
	}
}

func RandomIntegers(elements int) []int {
	rand.Seed(time.Now().UnixNano())
	var result []int = make([]int, 0)
	var min int = 0
	var max int = 1000000

	for i := 0; i < elements; i++ {
		result = append(result, rand.Intn(max-min)+min)
	}

	return result
}
