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

	areas, perimeters, _ := solve(input)
	for i := range areas {
		result += areas[i] * perimeters[i]
	}

	return result
}

func part2(fp string) int {
	input := loadInput(fp)
	result := 0

	areas, _, sides := solve(input)
	for i := range areas {
		result += areas[i] * sides[i]
	}

	return result
}

func solve(input [][]rune) (areas, perimeters, sides []int) {
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			s := input[y][x]
			if s > 'Z' {
				continue
			}

			mark := rune(s + 32)
			area, perimeter, side := traverse(input, s, mark, x, y)

			areas = append(areas, area)
			perimeters = append(perimeters, perimeter)
			sides = append(sides, side)
		}
	}

	return
}

func traverse(input [][]rune, symbol rune, mark rune, x, y int) (area, perimeter, sides int) {
	if y < 0 || x < 0 || y >= len(input) || x >= len(input[y]) {
		return 0, 1, 0
	}

	if input[y][x] != symbol {
		if input[y][x] == mark {
			return 0, 0, 0
		}

		return 0, 1, 0
	}

	input[y][x] = mark
	area = 1
	perimeter = 0

	dirs := [][2]int{
		{x, y - 1}, // north
		{x - 1, y}, // west
		{x, y + 1}, // south
		{x + 1, y}, // east
	}

	for _, dir := range dirs {
		a, p, s := traverse(input, symbol, mark, dir[0], dir[1])
		area += a
		perimeter += p
		sides += s
	}

	sides += countCorners(input, x, y)

	return
}

func countCorners(input [][]rune, x, y int) int {
	result := 0

	within := func(px, py int) bool {
		if px < 0 || py < 0 || py >= len(input) || px >= len(input[py]) {
			return false
		}

		return input[py][px] == input[y][x]
	}

	checks := [][3]bool{
		{within(x-1, y-1), within(x-1, y), within(x, y-1)}, // NW
		{within(x+1, y-1), within(x+1, y), within(x, y-1)}, // NE
		{within(x-1, y+1), within(x-1, y), within(x, y+1)}, // SW
		{within(x+1, y+1), within(x+1, y), within(x, y+1)}, // SE
	}

	for _, check := range checks {
		if check[1] != check[2] {
			continue
		}

		if check[0] == false {
			result++
		} else if check[1] == false && check[2] == false {
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
