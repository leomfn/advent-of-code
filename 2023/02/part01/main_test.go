package main

import "testing"

func TestGetSumOfGameIds(t *testing.T) {
	path := "input.example.txt"

	want := 8
	validGameSum := getSumOfGameIds(path)

	if validGameSum != want {
		t.Fatalf("expected %v, got %v", want, validGameSum)
	}
}
