package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{45, 45}, 90},
		test{[]int{4, 4, 4}, 12},
		test{[]int{1, 1, 1}, 3},
	}

	for _, v := range tests {
		x := Mysum(v.data...)
		if x != v.answer {
			t.Error("expecting", v.answer, "got", x)
		}
	}
}
