package main

import "testing"

func BenchmarkEscape(b *testing.B) {
	for range b.N {
		escape()
	}
}
