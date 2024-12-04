package main

import (
	"bufio"
	"os"
)

type Position struct {
	X int
	Y int
}

func main() {
	println("part1", part1("input.txt"))
	println("part2", part2("input.txt"))
}

func part1(fp string) int {
	result := 0

	input := loadInput(fp)

	for y := range input {
		for x := range input[y] {
			if input[y][x] != 'X' {
				continue
			}

			// north
			if y >= 3 &&
				input[y-1][x] == 'M' &&
				input[y-2][x] == 'A' &&
				input[y-3][x] == 'S' {
				result++
			}

			// northwest
			if y >= 3 && x >= 3 &&
				input[y-1][x-1] == 'M' &&
				input[y-2][x-2] == 'A' &&
				input[y-3][x-3] == 'S' {
				result++
			}

			// west
			if x >= 3 &&
				input[y][x-1] == 'M' &&
				input[y][x-2] == 'A' &&
				input[y][x-3] == 'S' {
				result++
			}

			// southwest
			if y < len(input)-3 && x >= 3 &&
				input[y+1][x-1] == 'M' &&
				input[y+2][x-2] == 'A' &&
				input[y+3][x-3] == 'S' {
				result++
			}

			// south
			if y < len(input)-3 &&
				input[y+1][x] == 'M' &&
				input[y+2][x] == 'A' &&
				input[y+3][x] == 'S' {
				result++
			}

			// southeast
			if y < len(input)-3 && x < len(input[y])-3 &&
				input[y+1][x+1] == 'M' &&
				input[y+2][x+2] == 'A' &&
				input[y+3][x+3] == 'S' {
				result++
			}

			// east
			if x < len(input[y])-3 &&
				input[y][x+1] == 'M' &&
				input[y][x+2] == 'A' &&
				input[y][x+3] == 'S' {
				result++
			}

			// northeast
			if y >= 3 && x < len(input[y])-3 &&
				input[y-1][x+1] == 'M' &&
				input[y-2][x+2] == 'A' &&
				input[y-3][x+3] == 'S' {
				result++
			}
		}
	}

	return result
}

func part2(fp string) int {
	result := 0

	input := loadInput(fp)

	validate := func(a, b rune) bool {
		return a != b && (a == 'M' || a == 'S') && (b == 'M' || b == 'S')
	}

	for y := 1; y < len(input)-1; y++ {
		for x := 1; x < len(input[y])-1; x++ {
			if input[y][x] != 'A' {
				continue
			}

			nw := input[y-1][x-1]
			sw := input[y+1][x-1]
			se := input[y+1][x+1]
			ne := input[y-1][x+1]

			if !validate(nw, se) || !validate(sw, ne) {
				continue
			}

			result++
		}
	}

	return result
}

func loadInput(fp string) [][]rune {
	file, _ := os.Open(fp)
	defer file.Close()

	input := make([][]rune, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, []rune(line))
	}

	return input
}
