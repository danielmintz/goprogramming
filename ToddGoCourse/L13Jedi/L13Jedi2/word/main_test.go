package word

import (
	"fmt"
	"goprogramming/ToddGoCourse/L13Jedi/L13Jedi2/quote"
	"testing"
)

func ExampleCount() {
	fmt.Println(Count("James Bond is alive"))
	// Output:
	// 4
}
func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func TestCount(t *testing.T) {
	s := Count("James Bond is alive")
	if s != 4 {
		t.Error("got", s, "want", 4)
	}

}
func TestUseCount(t *testing.T) {
	m := UseCount("James James is alive")
	for k, v := range m {
		switch k {
		case "James":
			if v != 2 {
				t.Error("got", v, "want", 2)
			}
		case "is":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "alive":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		}
	}
}
