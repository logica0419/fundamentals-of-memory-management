package main

type Data struct {
	value int
}

func main() {
	data := &Data{value: 42}
	_ = data.value
}
