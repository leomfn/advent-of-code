package main

import "testing"

func TestGetSumOfGearRatios(t *testing.T) {
	path := "input.example.txt"

	want := 467835
	result := getSumOfGearRatios(path)

	if result != want {
		t.Fatalf("expected %v, got %v", want, result)
	}
}
