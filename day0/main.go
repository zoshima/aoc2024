package main

import "os"

func main() {
	println("part1", part1("input.txt"))
	println("part2", part2("input.txt"))
}

func part1(fp string) int {
	return 0
}

func part2(fp string) int {
	return 0
}

func loadInput(fp string) any {
	file, _ := os.Open(fp)
	defer file.Close()

	return nil
}
