package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		test{[]int{2, 3, 4, 5, 6}, 4},
		test{[]int{1, 2, 3, 4, 5}, 3},
		test{[]int{5, 6, 7, 8, 9, 10, 11, 12}, 8.5},
	}
	for _, v := range tests {
		x := CenteredAvg(v.data)
		if x != v.answer {
			t.Error("got", x, "want", v.answer)
		}

	}
}
func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{2, 3, 4, 5, 6}))
	// Output:
	// 4
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}
