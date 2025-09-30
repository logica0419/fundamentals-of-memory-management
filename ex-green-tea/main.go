package main

import "runtime"

type T struct {
	a int
	b *int
}

func main() {
	s := make([]*T, 0)

	for i := range 1000000 {
		b := 10
		t := T{
			a: i,
			b: &b,
		}
		s = append(s, &t)
	}

	runtime.GC()

	_ = s
}
