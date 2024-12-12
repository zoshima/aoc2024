package main

import (
	"bufio"
	"os"
)

func main() {
	println("part1", part1("input.txt"))
	println("part2", part2("input.txt"))
}

func part1(fp string) int {
	input := loadInput(fp)
	result := 0

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			s := input[y][x]
			if s > 'Z' {
				continue
			}

			area, perimeter := traverse(input, s, x, y)
			result += area * perimeter
		}
	}

	return result
}

func part2(fp string) int {
	return 0
}

func traverse(input [][]rune, symbol rune, x, y int) (area, perimeter int) {
	if y < 0 || x < 0 || y >= len(input) || x >= len(input[y]) {
		return 0, 1
	}

	if input[y][x] != symbol {
		if input[y][x] == symbol+32 {
			return 0, 0
		}

		return 0, 1
	}

	input[y][x] = symbol + 32
	area = 1
	perimeter = 0

	dirs := [][2]int{
		{x, y - 1}, // north
		{x - 1, y}, // west
		{x, y + 1}, // south
		{x + 1, y}, // east
	}

	for _, dir := range dirs {
		a, p := traverse(input, symbol, dir[0], dir[1])
		area += a
		perimeter += p
	}

	return
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
