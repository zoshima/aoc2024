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
	return len(traverse(x, y, m))
}

func part2(fp string) int {
	x, y, m := loadInput(fp)
	result := 0

	for tile := range traverse(x, y, m) {
		cx := int(real(tile))
		cy := int(imag(tile))

		if cx == x && cy == y {
			continue
		}

		m[cy][cx] = '#'

		if traverse(x, y, m) == nil {
			result++
		}

		m[cy][cx] = '.'
	}

	return result
}

func traverse(x, y int, m [][]rune) map[complex64][]int {
	currentDirection := North
	traversedTiles := make(map[complex64][]int)

	for {
		key := complex(float32(x), float32(y))
		if directions, ok := traversedTiles[key]; ok {
			if slices.Contains(directions, currentDirection) {
				return nil // is looping
			}

			traversedTiles[key] = append(directions, currentDirection)
		} else {
			traversedTiles[key] = []int{currentDirection}
		}

		nx, ny := x, y

		switch currentDirection {
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
			if currentDirection == West {
				currentDirection = North
			} else {
				currentDirection++
			}

			continue
		}

		x, y = nx, ny
	}

	return traversedTiles
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
			}
		}

		lineNumber++
	}

	return
}
