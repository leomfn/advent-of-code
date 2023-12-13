package main

import "testing"

func TestGetNearestLocation(t *testing.T) {
	path := "input.example.txt"

	var want int = 35
	result := getNearestLocation(path)

	if result != want {
		t.Fatalf("expected %v, got %v", want, result)
	}
}
