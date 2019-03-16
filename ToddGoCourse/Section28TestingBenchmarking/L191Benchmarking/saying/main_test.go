package saying

import (
	"fmt"
	"testing"
)

func ExampleGreet() {
	fmt.Println(Greet("James"))
	// Output:
	// Welcome to the sayings package,James

}

func TestGreet(t *testing.T) {
	s := Greet("James")
	if s != "Welcome to the sayings package,James" {
		t.Error("got", s, "Expected", "Welcome to the sayings package,James")
	}

}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}
