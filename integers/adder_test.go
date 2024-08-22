package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("sum", func(t *testing.T) {
		sum := Add(1, 2)
		expected := 3
		if sum != expected {
			t.Errorf("Sum was incorrect, got: %d, expected: %d.", sum, expected)
		}
	})
}

func ExampleAdd() {
	sum := Add(4, 9)
	fmt.Println(sum)
	// Output: 13
}
