package main

func main() {
	num := 5
	slice := make([]int, 0)

	for i := range num {
		slice = append(slice, i)
	}

	_ = slice
}
