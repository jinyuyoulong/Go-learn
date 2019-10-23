package main

import (
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T) {
	b := &bytes.Buffer{}
	CountDown(b)
	got := b.String()
	// 反引号语法是创建 string 的另一种方式，但是允许你放置东西例如放到新的一行，对我们的测试来说是完美的。
	want := `3
2
1
GO!`
	if got != want {
		t.Errorf("got %s,but want %s", got, want)
	}
}
