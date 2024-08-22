package structs_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("Perimeter", func(t *testing.T) {
		rectangle := Rectangle{3.4, 5.5}
		got := Perimeter(rectangle)
		want := 17.8

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	validateResult := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2g want %.2g", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{3.8, 8.6}
		want := 32.68

		validateResult(t, rectangle, want)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793

		validateResult(t, circle, want)
	})

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}
}
