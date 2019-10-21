package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{width: 10.0, height: 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf("want '%.2f' but got '%.2f'", want, got)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("want '%.2f' but got '%.2f'", want, got)
		}
	}
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{width: 10.0, height: 10.0}

		want := 100.0
		checkArea(t, rectangle, want)

	})
	t.Run("cycle", func(t *testing.T) {
		cycle := Cycle{10.0}
		want := 314.1592653589793
		checkArea(t, cycle, want)
	})
}
