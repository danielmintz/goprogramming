package main

import "testing"

func TestMySum(t *testing.T) {
	x := MySum(1, 2)
	if x != 3 {
		t.Error("expecting 3", "got", x)
	}
}
