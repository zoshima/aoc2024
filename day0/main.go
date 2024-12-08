package main

import "os"

func main() {
	input := loadInput("input.txt")
	println("part1", part1(input))
	println("part2", part2(input))
}

func part1(input any) int {
	return 0
}

func part2(input any) int {
	return 0
}

func loadInput(fp string) any {
	file, _ := os.Open(fp)
	defer file.Close()

	return nil
}
