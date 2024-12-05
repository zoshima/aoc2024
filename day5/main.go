package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	println("part1", part1("input.txt"))
	println("part2", part2("input.txt"))
}

// map[29:[13] 47:[53 13 61 29] 53:[29 13] 61:[13 53 29] 75:[29 53 47 61 13] 97:[13 61 47 29 53 75]]
// [[75 47 61 53 29] [97 61 53 29 13] [75 29 13] [75 97 47 61 53] [61 13 29] [97 13 75 29 47]]

func part1(fp string) int {
	m, rows := loadInput(fp)
	result := 0

RangeRows:
	for _, row := range rows {
		for i, n := range row {
			entry, ok := m[n]
			if !ok {
				continue
			}

			for j := 0; j < i; j++ {
				if slices.Contains(entry, row[j]) {
					continue RangeRows
				}
			}
		}

		mi := len(row) / 2
		result += row[mi]
	}

	return result
}

func part2(fp string) int {
	return 0
}

func loadInput(fp string) (map[int][]int, [][]int) {
	file, _ := os.Open(fp)
	defer file.Close()

	m := make(map[int][]int)
	a := make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		values := strings.Split(line, "|")
		left, _ := strconv.Atoi(values[0])
		right, _ := strconv.Atoi(values[1])

		if v, ok := m[left]; ok {
			m[left] = append(v, right)
		} else {
			m[left] = []int{right}
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		values := strings.Split(line, ",")
		row := make([]int, len(values))
		for i, value := range values {
			row[i], _ = strconv.Atoi(value)
		}

		a = append(a, row)
	}

	return m, a
}
