package main

import (
	"bufio"
	"os"
	"slices"
)

const (
	North = iota
	East
	South
	West
)

func main() {
	println("part1", part1("input.txt"))
	println("part2", part2("input.txt"))
}

func part1(fp string) int {
	x, y, m := loadInput(fp)
	direction := North
	result := 0

	for {
		if m[y][x] == '.' {
			m[y][x] = 'X'
			result++
		}

		nx, ny := x, y

		switch direction {
		case North:
			ny -= 1
		case East:
			nx += 1
		case South:
			ny += 1
		case West:
			nx -= 1
		}

		if ny < 0 || nx < 0 || ny >= len(m) || nx >= len(m[y]) {
			break
		}

		if m[ny][nx] == '#' {
			direction++
			if direction > West {
				direction = North
			}
		} else {
			x, y = nx, ny
		}
	}

	return result
}

func part2(fp string) int {
	return 0
}

func loadInput(fp string) (x int, y int, m [][]rune) {
	file, _ := os.Open(fp)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		line := []rune(scanner.Text())
		m = append(m, line)

		if x == 0 && y == 0 {
			colNumber := slices.Index(line, '^')
			if colNumber != -1 {
				x, y = colNumber, lineNumber
				m[y][x] = '.'
			}
		}

		lineNumber++
	}

	return
}
