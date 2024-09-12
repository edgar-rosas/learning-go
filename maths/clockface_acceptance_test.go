package maths

import (
	"math"
	"testing"
	"time"
)

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(t, 0, 0, 30), Point{0, -1}},
		{simpleTime(t, 0, 0, 45), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(t, c.time), func(t *testing.T) {
			got := secondHandPoint(c.time)
			if !roughlyEqualPoint(t, got, c.point) {
				t.Fatalf("Wanted %v Point, but got %v", c.point, got)
			}
		})
	}
}

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := Point{X: 150, Y: 150 - 90}
	got := SecondHand(tm)

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	want := Point{X: 150, Y: 150 + 90}
	got := SecondHand(tm)

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}

func roughlyEqualFloat64(t testing.TB, a, b float64) bool {
	t.Helper()
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(t testing.TB, a, b Point) bool {
	t.Helper()
	return roughlyEqualFloat64(t, a.X, b.X) &&
		roughlyEqualFloat64(t, a.Y, b.Y)
}
