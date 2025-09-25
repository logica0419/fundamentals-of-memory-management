//go:build ignore

package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 0)
	num := 5

	for i := range num {
		slice = append(slice, i)
	}

	fmt.Println(slice)
}
