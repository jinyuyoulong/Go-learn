package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0
	if got != want {
		t.Errorf("want '%.2f' but got '%.2f'", want, got)
	}
}

func TestArea(t *testing.T) {
	got := Area(10.0, 10.0)
	want := 100.0
	if got != want {
		t.Errorf("want '%.2f' but got '%.2f'", want, got)
	}
}
