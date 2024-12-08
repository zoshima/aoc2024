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

	antennas := make(map[rune][][2]int)
	for y, line := range input {
		for x, r := range line {
			if r == '.' {
				continue
			}

			if positions, ok := antennas[r]; ok {
				antennas[r] = append(positions, [2]int{y, x})
			} else {
				antennas[r] = [][2]int{{y, x}}
			}
		}
	}

	setAntinode := func(y, x int) bool {
		if y >= 0 && y < len(input) && x >= 0 && x < len(input[y]) {
			if input[y][x] != '#' {
				input[y][x] = '#'
				return true
			}
		}

		return false
	}

	for _, positions := range antennas {
		for i := 0; i < len(positions)-1; i++ {
			for j := i + 1; j < len(positions); j++ {
				y1, x1 := positions[i][0], positions[i][1]
				y2, x2 := positions[j][0], positions[j][1]

				dy, dx := y2-y1, x2-x1

				y, x := y1-dy, x1-dx
				if setAntinode(y, x) {
					result++
				}

				y, x = y2+dy, x2+dx
				if setAntinode(y, x) {
					result++
				}
			}
		}
	}

	return result
}

func part2(input any) int {
	return 0
}

func loadInput(fp string) [][]rune {
	file, _ := os.Open(fp)
	defer file.Close()

	m := make([][]rune, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		m = append(m, []rune(line))
	}

	return m
}
