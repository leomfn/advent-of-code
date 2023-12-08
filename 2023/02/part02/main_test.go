package main

import "testing"

func TestGetSumOfPowers(t *testing.T) {
	path := "input.example.txt"

	want := 2286
	validGameSum := getSumOfPowers(path)

	if validGameSum != want {
		t.Fatalf("expected %v, got %v", want, validGameSum)
	}
}
