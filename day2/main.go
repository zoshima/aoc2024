package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	println("part1", part1("input.txt"))
}

func part1(fp string) int {
	input := loadInput(fp)
	result := 0

RowLoop:
	for _, row := range input {
		var wantDecreasing any

		for i := 1; i < len(row); i++ {
			delta := row[i-1] - row[i]
			if delta == 0 {
				continue RowLoop
			}

			isDecreasing := delta > 0
			if !isDecreasing {
				delta = -delta
			}

			if delta > 3 {
				continue RowLoop
			}

			if wantDecreasing == nil {
				wantDecreasing = isDecreasing
			} else if wantDecreasing != isDecreasing {
				continue RowLoop
			}
		}

		result++
	}

	return result
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
