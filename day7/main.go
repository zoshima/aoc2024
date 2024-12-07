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
		if validate(eq, 1, eq.Operands[0], false) {
			result += eq.Result
		}
	}

	return result
}

func part2(fp string) int {
	input := loadInput(fp)
	var result int

	for i, eq := range input {
		if validate(eq, 1, eq.Operands[0], true) {
			println(i)
			result += eq.Result
		}
	}

	return result
}

func validate(eq *Equation, index int, currentOperand int, ext bool) bool {
	if index == len(eq.Operands) {
		return currentOperand == eq.Result
	}

	if currentOperand > eq.Result {
		return false
	}

	operands := [3]int{
		currentOperand + eq.Operands[index],
		currentOperand * eq.Operands[index],
	}

	if ext {
		operands[2], _ = strconv.Atoi(strconv.Itoa(currentOperand) + strconv.Itoa(eq.Operands[index]))
	}

	for _, operand := range operands {
		if validate(eq, index+1, operand, ext) {
			return true
		}
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
