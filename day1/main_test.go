package main

import "testing"

func TestPrint(t *testing.T) {
	exp := "day1"

	res := print()

	if res != exp {
		t.Fatalf("tests[%d] - exp: %q, got %q", 1, exp, res)
	}
}
