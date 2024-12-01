package main

import (
	"bufio"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func main() {
	println("part1", part1("input.txt"))
	println("part2", part2("input.txt"))
}

func part1(input string) int {
	left, right := loadInput(input)

	slices.Sort(left)
	slices.Sort(right)

	result := 0
	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		if diff < 0 {
			diff = -diff
		}

		result += diff
	}

	return result
}

func part2(input string) int {
	left, right := loadInput(input)

	m := make(map[int]int)
	for _, n := range right {
		if _, ok := m[n]; ok {
			m[n] += 1
		} else {
			m[n] = 1
		}
	}

	result := 0
	for _, n := range left {
		if count, ok := m[n]; ok {
			result += n * count
		}
	}

	return result
}

func loadInput(path string) ([]int, []int) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	regex := regexp.MustCompile("[0-9]+")

	left := make([]int, 0)
	right := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		matches := regex.FindAll(line, -1)

		l, _ := strconv.Atoi(string(matches[0]))
		r, _ := strconv.Atoi(string(matches[1]))

		left = append(left, l)
		right = append(right, r)
	}

	return left, right
}
