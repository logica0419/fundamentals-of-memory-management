//go:build ignore

package main

func main() {
	slice := createSlice()

	for _, v := range slice {
		println(v)
	}
}

func createSlice() []int {
	slice := make([]int, 0)
	num := 5

	for i := range num {
		slice = append(slice, i)
	}

	return slice
}
