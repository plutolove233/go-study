package go122_test

import (
	"math/rand"
	rand2 "math/rand/v2"
	"testing"
)

func BenchmarkRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Intn(100)
	}
}

func BenchmarkRandv2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand2.IntN(100)
	}
}