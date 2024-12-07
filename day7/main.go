package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type Equation struct {
	Result   int
	Operands []int
}

func main() {
	println("part1", part1("input.txt"))
	println("part2", part2("input.txt"))
}

func part1(fp string) int {
	input := loadInput(fp)
	var result int

	for _, eq := range input {
		if validate(eq.Operands, eq.Result) {
			result += eq.Result
		}
	}

	return result
}

func part2(fp string) int {
	return 0
}

func validate(arr []int, lim int) bool {
	if len(arr) < 2 {
		return false
	}

	sum := arr[0] + arr[1]
	if sum == lim {
		return true
	}

	prod := arr[0] * arr[1]
	if prod == lim {
		return true
	}

	if sum < lim && validate(append([]int{sum}, arr[2:]...), lim) {
		return true
	}

	if prod < lim && validate(append([]int{prod}, arr[2:]...), lim) {
		return true
	}

	return false
}

func loadInput(fp string) []*Equation {
	file, _ := os.Open(fp)
	defer file.Close()

	input := make([]*Equation, 0)
	regex := regexp.MustCompile("[0-9]+")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		matches := regex.FindAll(line, -1)
		eq := Equation{}

		eq.Result, _ = strconv.Atoi(string(matches[0]))
		eq.Operands = make([]int, len(matches)-1)

		for i := 0; i < len(eq.Operands); i++ {
			eq.Operands[i], _ = strconv.Atoi(string(matches[i+1]))
		}

		input = append(input, &eq)
	}

	return input
}
