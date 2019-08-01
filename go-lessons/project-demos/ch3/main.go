package main

import (
	"fmt"
)

func main() {
	var data [32]byte = [32]byte{1, 2, 3}
	zero(&data)
	fmt.Println(data)
}

func zero(ptr *[32]byte) {
	// for i := range ptr {
	// 	ptr[i] = 0
	// }
	*ptr = [32]byte{}
}
