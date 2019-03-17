package dog

import "testing"

func TestYears(t *testing.T) {
	s := Years(1)
	if s != 7 {
		t.Error("Got", s, "Want", 7)
	}
}

func TestYearsTwo(t *testing.T) {
	s := YearsTwo(1)
	if s != 7 {
		t.Error("Got", s, "Want", 7)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(1)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(1)
	}
}
