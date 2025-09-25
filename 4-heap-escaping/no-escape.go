//go:build ignore

package main

func main() {
	slice := make([]int, 0)
	num := 5

	for i := range num {
		slice = append(slice, i)
	}

	for _, v := range slice {
		println(v)
	}
}

