package main

import "testing"

func TestMysum(t *testing.T) {
	x := mySum(2, 3)
	if x != 5 {
		t.Error("expected", 5, "got", x)

	}

}
