package main

import "testing"

func TestPart1(t *testing.T) {
	want := 55312
	got := part1([]int{125, 17})

	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	want := 0
	got := part2([]int{125, 17})

	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}
}
