package main

import "strconv"

type Stone struct {
	Next        *Stone
	Value       int
	StringValue string
}

func (s *Stone) SetValue(v int) {
	s.Value = v
	s.StringValue = strconv.Itoa(v)
}

func main() {
	input := []int{4610211, 4, 0, 59, 3907, 201586, 929, 33750}

	println("part1", part1(input))
	println("part2", part2(input))
}

func part1(input []int) int {
	result := len(input)

	first := parseInput(input)
	for i := 0; i < 25; i++ {
		stone := first
		for stone != nil {
			if stone.Value == 0 {
				stone.SetValue(1)
			} else if len(stone.StringValue)%2 == 0 {
				splitIndex := len(stone.StringValue) / 2

				leftValue, _ := strconv.Atoi(stone.StringValue[:splitIndex])
				rightValue, _ := strconv.Atoi(stone.StringValue[splitIndex:])

				ns := Stone{}
				ns.SetValue(rightValue)
				ns.Next = stone.Next

				stone.SetValue(leftValue)
				stone.Next = &ns
				stone = &ns

				result++
			} else {
				stone.SetValue(stone.Value * 2024)
			}

			stone = stone.Next
		}
	}

	return result
}

func part2(input []int) int {
	return 0
}

func parseInput(input []int) *Stone {
	first := Stone{}
	first.SetValue(input[0])

	prev := &first

	for i := 1; i < len(input); i++ {
		stone := Stone{}
		stone.SetValue(input[i])

		prev.Next = &stone
		prev = &stone
	}

	return &first
}
