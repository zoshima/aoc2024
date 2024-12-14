package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type Machine struct {
	A     [2]int
	B     [2]int
	Prize [2]int
}

func main() {
	println("part1", part1("input.txt"))
	println("part2", part2("input.txt"))
}

func part1(fp string) int {
	// input := loadInput(fp)

	return 0
}

func part2(fp string) int {
	return 0
}

func loadInput(fp string) []*Machine {
	file, _ := os.Open(fp)
	defer file.Close()

	regex := regexp.MustCompile("[0-9]+")

	machines := make([]*Machine, 0)
	machine := &Machine{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			machines = append(machines, machine)
			machine = &Machine{}
			continue
		}

		matches := regex.FindAll(line, -1)
		left, _ := strconv.Atoi(string(matches[0]))
		right, _ := strconv.Atoi(string(matches[1]))

		if machine.A[0] == 0 && machine.A[1] == 0 {
			machine.A[0], machine.A[1] = left, right
			continue
		}

		if machine.B[0] == 0 && machine.B[1] == 0 {
			machine.B[0], machine.B[1] = left, right
			continue
		}

		if machine.Prize[0] == 0 && machine.Prize[1] == 0 {
			machine.Prize[0], machine.Prize[1] = left, right
			continue
		}
	}

	return machines
}
