package main

import (
	"testing"
)

func BenchmarkCubeConundrum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res, err := cubeConundrum()
		if err != nil {
			b.Fatal(err)
		}

		if res != 2076 {
			b.Errorf("Failed - expected %d, got %d", 2076, res)
		}
	}
}
