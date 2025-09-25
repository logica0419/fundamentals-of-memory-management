package main

func noEscapeAlloc() {
	num := 5
	slice := make([]int, 0)

	for i := range num {
		slice = append(slice, i)
	}

	for _, v := range slice {
		println(v)
	}
}
