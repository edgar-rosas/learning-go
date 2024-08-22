package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat letter 6 times", func(t *testing.T) {
		repeated := Repeat("a", 6)
		expected := "aaaaaa"
		if repeated != expected {
			t.Errorf("repeat result was incorrect, got: %s, expected: %s.", repeated, expected)
		}
	})
}

func ExampleRepeat() {
	repeated := Repeat("a", 6)
	fmt.Println(repeated)
	// Output: aaaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}
