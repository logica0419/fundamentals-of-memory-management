//go:build ignore

package main

import (
	"fmt"
)

func main() {
	num := 5
	f(num)

	fmt.Println(num)
}

func f(num int) {
	num = 10
}
