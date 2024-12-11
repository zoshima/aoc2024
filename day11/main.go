package main

import "math"

type Stone struct {
	Next  *Stone
	Value int
}

func main() {
	input := []int{4610211, 4, 0, 59, 3907, 201586, 929, 33750}

	println("part1", part1(input))
	println("part2", part2(input))
}

func part1(input []int) int {
	m := make(map[int]int)
	for i := range input {
		m[input[i]] = 1
	}

	for i := 0; i < 25; i++ {
		m = blink(m)
	}

	result := 0
	for _, v := range m {
		result += v
	}

	return result
}

func part2(input []int) int {
	m := make(map[int]int)
	for i := range input {
		m[input[i]] = 1
	}

	for i := 0; i < 75; i++ {
		m = blink(m)
	}

	result := 0
	for _, v := range m {
		result += v
	}

	return result
}

func blink(stones map[int]int) map[int]int {
	m := make(map[int]int)

	for value, count := range stones {
		if value == 0 {
			increment(m, 1, count)
			continue
		}

		valueLength := countDigits(value)
		if valueLength%2 == 0 {
			left, right := splitValue(value, valueLength)
			increment(m, left, count)
			increment(m, right, count)

			continue
		}

		increment(m, value*2024, count)
	}

	return m
}

func increment(m map[int]int, key int, value int) {
	if _, ok := m[key]; ok {
		m[key] += value
	} else {
		m[key] = value
	}
}

func splitValue(value int, numDigits int) (int, int) {
	divisor := int(math.Pow10(numDigits / 2))

	return value / divisor, value % divisor
}

func countDigits(value int) int {
	if value < 10 {
		return 1
	}

	c := 0
	for value != 0 {
		value /= 10
		c++
	}

	return c
}
