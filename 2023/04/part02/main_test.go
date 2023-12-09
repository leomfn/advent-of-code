package main

import "testing"

func TestGetCountOfScratchCards(t *testing.T) {
	path := "input.example.txt"

	var want int = 30
	result := getCountOfScratchCards(path)

	if result != want {
		t.Fatalf("expected %v, got %v", want, result)
	}
}
