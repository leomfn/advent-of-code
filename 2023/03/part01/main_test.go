package main

import "testing"

func TestGetSumOfPartNumbers(t *testing.T) {
	path := "input.example.txt"

	want := 4361
	result := getSumOfPartNumbers(path)

	if result != want {
		t.Fatalf("expected %v, got %v", want, result)
	}
}
