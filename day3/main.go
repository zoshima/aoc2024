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

	for _, instruction := range input {
		if instruction[:3] == "mul" {
			result += eval(instruction)
		}
	}

	return result
}

func part2(fp string) int {
	input := loadInput(fp)
	result := 0

	do := true

	for _, instruction := range input {
		fn := instruction[:3]

		switch fn {
		case "mul":
			if !do {
				continue
			}

			result += eval(instruction)
		case "don":
			do = false
		default:
			do = true
		}
	}

	return result
}

func eval(instruction string) int {
	i := strings.Index(instruction, ",")

	multiplicand, _ := strconv.Atoi(instruction[4:i])
	multiplier, _ := strconv.Atoi(instruction[i+1 : len(instruction)-1])

	return multiplicand * multiplier
}

func loadInput(fp string) []string {
	input := make([]string, 0)
	regex := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)|don't\(\)|do\(\)`)

	file, _ := os.Open(fp)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		matches := regex.FindAllString(line, -1)
		input = append(input, matches...)
	}

	return input
}
