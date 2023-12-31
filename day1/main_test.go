package main

import (
	"testing"
)

func BenchmarkTrebuchet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := trebuchet()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestTrebuchet(t *testing.T) {
	_, err := trebuchet()
	if err != nil {
		t.Fatal(err)
	}
}
