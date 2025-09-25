package main

func noEscapeNoAlloc() {
	num := 5
	slice := make([]int, 0, num)

	for i := range num {
		slice = append(slice, i)
	}

	for _, v := range slice {
		println(v)
	}
}
