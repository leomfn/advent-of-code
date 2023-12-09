package main

import "testing"

func TestGetSumOfPoints(t *testing.T) {
	path := "input.example.txt"

	var want uint = 13
	result := getSumOfPoints(path)

	if result != want {
		t.Fatalf("expected %v, got %v", want, result)
	}
}
