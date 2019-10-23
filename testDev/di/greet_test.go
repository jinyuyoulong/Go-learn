package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	b := bytes.Buffer{}
	Greet(&b, "fan")
	got := b.String()
	want := "hello, fan"
	if got != want {
		t.Errorf("got %s,but want %s", got, want)
	}
}
