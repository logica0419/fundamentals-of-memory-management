package main

import "testing"

func BenchmarkNoEscapeAlloc(b *testing.B) {
	for range b.N {
		noEscapeAlloc()
	}
}
