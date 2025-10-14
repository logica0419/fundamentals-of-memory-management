package main

import "fmt"

func recursive(n int) {
	n++
	fmt.Println(n)
	recursive(n)
}

func main() {
	recursive(0)
}
