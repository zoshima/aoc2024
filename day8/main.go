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
	antennas := getAntennas(input)
	result := 0

	for _, positions := range antennas {
		for i := 0; i < len(positions)-1; i++ {
			for j := i + 1; j < len(positions); j++ {
				y1, x1 := positions[i][0], positions[i][1]
				y2, x2 := positions[j][0], positions[j][1]

				dy, dx := y2-y1, x2-x1

				y, x := y1-dy, x1-dx
				if setAntinode(input, y, x) == 1 {
					result++
				}

				y, x = y2+dy, x2+dx
				if setAntinode(input, y, x) == 1 {
					result++
				}
			}
		}
	}

	return result
}

func part2(fp string) int {
	input := loadInput(fp)
	antennas := getAntennas(input)
	result := 0

	for _, positions := range antennas {
		for i := 0; i < len(positions)-1; i++ {
			for j := i + 1; j < len(positions); j++ {
				y1, x1 := positions[i][0], positions[i][1]
				y2, x2 := positions[j][0], positions[j][1]

				dy, dx := y2-y1, x2-x1

				y, x := y1, x1
				for {
					exitCode := setAntinode(input, y, x)
					if exitCode == 0 {
						break
					} else if exitCode == 1 {
						result++
					}

					y, x = y-dy, x-dx
				}

				y, x = y2, x2
				for {
					exitCode := setAntinode(input, y, x)
					if exitCode == 0 {
						break
					} else if exitCode == 1 {
						result++
					}

					y, x = y+dy, x+dx
				}
			}
		}
	}

	return result
}

func getAntennas(m [][]rune) map[rune][][2]int {
	antennas := make(map[rune][][2]int)
	for y, line := range m {
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

	return antennas
}

func setAntinode(m [][]rune, x, y int) int {
	if y >= 0 && y < len(m) && x >= 0 && x < len(m[y]) {
		if m[y][x] != '#' {
			m[y][x] = '#'
			return 1
		}

		return 2
	}

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
