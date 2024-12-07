package main

import "testing"

func TestPart1(t *testing.T) {
	want := 3749
	got := part1("input_test.txt")

	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	want := 11387
	got := part2("input_test.txt")

	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}
}
