package main

import (
	"bytes"
	"fmt"
)

func Greet(b *bytes.Buffer, name string) {
	fmt.Fprintf(b, "hello, %s", name)
}
