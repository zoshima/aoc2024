package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	println("part1", part1("input.txt"))
	println("part2", part2("input.txt"))
}

func part1(fp string) int {
	input := loadInput(fp)
	result := 0

	for _, line := range input {
		result += eval(line)
	}

	return result
}

func part2(fp string) int {
	return 0
}

func eval(instruction string) int {
	i := strings.Index(instruction, ",")

	multiplicand, _ := strconv.Atoi(instruction[4:i])
	multiplier, _ := strconv.Atoi(instruction[i+1 : len(instruction)-1])

	return multiplicand * multiplier
}

func loadInput(fp string) []string {
	input := make([]string, 0)
	regex := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)

	file, _ := os.Open(fp)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		matches := regex.FindAllString(line, -1)
		input = append(input, matches...)
	}

	return input
}
