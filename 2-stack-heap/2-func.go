//go:build ignore

package main

import "fmt"

func main() {
	v := 5
	v = a(v)

	fmt.Println(v)
}

func a(arg int) int {
	a := arg + 2
	return b(a)
}

func b(arg int) int {
	b := 1
	return arg + b
}
