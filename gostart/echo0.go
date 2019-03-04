package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// s := "你好世界"
	// const nativeString = `this is a go native string value`
	// s1 := "世界"
	// s2 := "\xe4\xb8\x96\xe7\x95\x8c"
	// s3 := "\u4e16\u754c"
	// s4 := "\U00004e16\U0000754c"
	s := "hello 世界"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

}
