package main

import "testing"

func BenchmarkGetRandomWord10(b *testing.B) {
	// run the getRandomWord function b.N times
	for n := 0; n < b.N; n++ {
		GetRandomWord()
	}
}
