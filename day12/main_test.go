package main

import "testing"

func TestPart1(t *testing.T) {
	want := 1930
	got := part1("input_test.txt")

	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	want := 1206
	got := part2("input_test.txt")

	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}

	want = 236
	got = part2("input_test_2.txt")

	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}

	want = 368
	got = part2("input_test_3.txt")

	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}
}
