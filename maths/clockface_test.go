package maths

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(t, 0, 0, 30), math.Pi},
		{simpleTime(t, 0, 0, 0), 0},
		{simpleTime(t, 0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(t, 0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, c := range cases {
		t.Run(testName(t, c.time), func(t *testing.T) {
			got := secondsInRadians(c.time)
			if got != c.angle {
				t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func simpleTime(t testing.TB, hours, minutes, seconds int) time.Time {
	t.Helper()
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(te testing.TB, t time.Time) string {
	te.Helper()
	return t.Format("15:04:05")
}
