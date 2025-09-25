package main

import "testing"

func BenchmarkNoEscapeNoAlloc(b *testing.B) {
	for range b.N {
		noEscapeNoAlloc()
	}
}
