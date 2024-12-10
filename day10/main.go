package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	println("part1", part1("input.txt"))
	println("part2", part2("input.txt"))
}

func part1(fp string) int {
	result := 0

	input := loadInput(fp)
	for y := range input {
		for x := range input[y] {
			if input[y][x] != 0 {
				continue
			}

			peaks := make(map[string]int)
			traverse(input, peaks, x, y, 0)
			result += len(peaks)
		}
	}

	return result
}

func part2(fp string) int {
	result := 0

	input := loadInput(fp)
	for y := range input {
		for x := range input[y] {
			if input[y][x] != 0 {
				continue
			}

			peaks := make(map[string]int)
			traverse(input, peaks, x, y, 0)

			for _, v := range peaks {
				result += v
			}
		}
	}

	return result
}

func traverse(m [][]int, peaks map[string]int, x, y, height int) {
	if y < 0 || y >= len(m) || x < 0 || x >= len(m[y]) {
		return
	}

	if m[y][x] != height {
		return
	}

	if height == 9 {
		key := fmt.Sprintf("%d:%d", y, x)
		if n, ok := peaks[key]; ok {
			peaks[key] = n + 1
		} else {
			peaks[key] = 1
		}

		return
	}

	traverse(m, peaks, x, y-1, height+1) // north
	traverse(m, peaks, x-1, y, height+1) // west
	traverse(m, peaks, x, y+1, height+1) // south
	traverse(m, peaks, x+1, y, height+1) // east
}

func loadInput(fp string) [][]int {
	file, _ := os.Open(fp)
	defer file.Close()

	input := make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()

		row := make([]int, len(line))
		for i := range line {
			row[i] = int(line[i] - '0')
		}

		input = append(input, row)
	}

	return input
}
