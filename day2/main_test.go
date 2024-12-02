package main

import "testing"

func TestPart1(t *testing.T) {
	want := 2
	got := part1("input_test.txt")

	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	want := 4
	got := part2("input_test.txt")

	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}
}
