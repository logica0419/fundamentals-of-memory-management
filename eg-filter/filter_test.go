package main

import "testing"

func BenchmarkFilterBefore(b *testing.B) {
	var list targetList = make([]int, 100)
	for i := 0; i < 100; i++ {
		list[i] = i
	}

	for b.Loop() {
		_ = list.FilterBefore(func(i int) bool {
			return i%2 == 0
		})
	}
}

func BenchmarkFilterAfter(b *testing.B) {
	var list targetList = make([]int, 100)
	for i := 0; i < 100; i++ {
		list[i] = i
	}

	for b.Loop() {
		_ = list.FilterAfter(func(i int) bool {
			return i%2 == 0
		})
	}
}
