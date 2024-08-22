package arrays

import (
	"fmt"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum of 5 integers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	fmt.Println(got)
	// Output: 15
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1, 2, 3, 4, 5})
	}
}

func TestSumAll(t *testing.T) {
	t.Run("sum of all integers", func(t *testing.T) {
		got := SumAll([]int{3, 9}, []int{18, 9})
		want := []int{12, 27}

		if !slices.Equal(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func ExampleSumAll() {
	got := SumAll([]int{3, 9}, []int{18, 9})
	fmt.Println(got)
	// Output: [12 27]
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{3, 9}, []int{18, 9})
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		if !slices.Equal(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	}
	t.Run("sum of all tails", func(t *testing.T) {
		got := SumAllTails([]int{3, 9, 5}, []int{18, 9, 2})
		want := []int{14, 11}
		checkSums(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{18, 9, 2})
		want := []int{0, 11}
		checkSums(t, got, want)
	})
}

func ExampleSumAllTails() {
	got := SumAllTails([]int{3, 9, 5}, []int{18, 9, 2})
	fmt.Println(got)
	// Output: [14 11]
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails([]int{3, 9, 5}, []int{18, 9, 2})
	}
}
