package main

func escape() {
	slice := createSlice()

	for _, v := range slice {
		println(v)
	}
}

func createSlice() []int {
	num := 5
	slice := make([]int, 0, 5)

	for i := range num {
		slice = append(slice, i)
	}

	return slice
}
