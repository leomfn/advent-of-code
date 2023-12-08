package main

import (
	"testing"
)

func TestTotalSum(t *testing.T) {
	dataPath := "input.example.txt"
	want := 281
	sum := totalSum(dataPath)

	if want != sum {
		t.Fatalf(`Expected %v, but got %v`, want, sum)
	}

}
