//go:build ignore

package main

import (
	"fmt"
)

func main() {
	num := 5
	change(num)

	fmt.Println(num)
}

func change(num int) {
	num = 10
}
