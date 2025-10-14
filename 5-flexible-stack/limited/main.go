package main

import (
	"fmt"
	"runtime/debug"
)

func recursive(n int) {
	n++
	fmt.Println(n)
	recursive(n)
}

func main() {
	debug.SetMaxStack(8000000)
	recursive(0)
}
