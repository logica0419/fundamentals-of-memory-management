package main

import "testing"

func BenchmarkMain(b *testing.B) {
	for b.Loop() {
		main()
	}
}
