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

func part1(fp string) int {
	m, rows := loadInput(fp)
	result := 0

	for _, row := range rows {
		i, _ := isCorrect(m, row)
		if i != -1 {
			continue
		}

		mi := len(row) / 2
		result += row[mi]
	}

	return result
}

func part2(fp string) int {
	m, rows := loadInput(fp)
	result := 0

	for _, row := range rows {
		index, conflict := isCorrect(m, row)
		if index == -1 {
			continue
		}

		for index != -1 {
			row[index], row[conflict] = row[conflict], row[index]
			index, conflict = isCorrect(m, row)
		}

		mi := len(row) / 2
		result += row[mi]
	}

	return result
}

func isCorrect(m map[int][]int, row []int) (index int, conflict int) {
	for i, n := range row {
		entry, ok := m[n]
		if !ok {
			continue
		}

		for j := 0; j < i; j++ {
			if slices.Contains(entry, row[j]) {
				return i, j
			}
		}
	}

	return -1, -1
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
