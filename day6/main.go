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

type Position struct {
	x int
	y int
}

func part1(fp string) int {
	x, y, m := loadInput(fp)
	positions, _ := traverse(x, y, North, m)
	return len(positions)
}

func part2(fp string) int {
	x, y, m := loadInput(fp)
	result := 0

	positions, _ := traverse(x, y, North, m)
	for position := range positions {
		if position.x == x && position.y == y {
			continue
		}

		m[position.y][position.x] = '#'
		_, isLooping := traverse(x, y, North, m)
		if isLooping {
			result++
		}

		m[position.y][position.x] = '.'
	}

	return result
}

func traverse(x, y, direction int, m [][]rune) (map[Position][]int, bool) {
	locations := make(map[Position][]int)
	isLooping := false

	for {
		key := Position{x, y}
		if directions, ok := locations[key]; ok {
			if slices.Contains(directions, direction) {
				isLooping = true
				break
			} else {
				locations[key] = append(directions, direction)
			}
		} else {
			locations[key] = []int{direction}
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

	return locations, isLooping
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
