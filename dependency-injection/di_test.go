package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Edgar")

	got := buffer.String()
	want := "Hello, Edgar"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
