package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	println("part1", part1("input.txt"))
	println("part2", part2("input.txt"))
}

func part1(fp string) int {
	input := loadInput(fp)
	result := 0

	for _, row := range input {
		if !isValid(row, -1) {
			continue
		}

		result++
	}

	return result
}

func part2(fp string) int {
	input := loadInput(fp)
	result := 0

RangeInput:
	for _, row := range input {
		if isValid(row, -1) {
			result++
			continue
		}

		for i := range row {
			if isValid(row, i) {
				result++
				continue RangeInput
			}
		}
	}

	return result
}

func isValid(row []int, skipIndex int) bool {
	var wantDecreasing any

	for i := 0; i < len(row)-1; i++ {
		if i == skipIndex {
			continue
		}

		j := i + 1
		if j == skipIndex {
			j++
		}

		if j == len(row) {
			break
		}

		delta := row[i] - row[j]
		if delta == 0 || delta < -3 || delta > 3 {
			return false
		}

		isDecreasing := delta > 0
		if wantDecreasing == nil {
			wantDecreasing = isDecreasing
		} else if wantDecreasing != isDecreasing {
			return false
		}
	}

	return true
}

func loadInput(fp string) [][]int {
	file, _ := os.Open(fp)
	scanner := bufio.NewScanner(file)

	input := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)

		row := make([]int, len(values))
		for i, value := range values {
			row[i], _ = strconv.Atoi(value)
		}

		input = append(input, row)
	}

	return input
}
