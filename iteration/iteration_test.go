package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	expected := "aaaaa"
	if got != expected {
		t.Errorf("got %s expected %s", got, expected)
	}
}

func ExampleRepeat() {
	got := Repeat("a")
	fmt.Println(got)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
