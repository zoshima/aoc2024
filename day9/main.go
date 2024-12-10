package main

import "os"

type File struct {
	ID   int
	Size int
}

func main() {
	println("part1", part1("input.txt"))
	println("part2", part2("input.txt"))
}

func part1(fp string) int {
	result := 0
	input := loadInput(fp)

	j := len(input) - 1
	for i := 0; i < j; i++ {
		if input[i] == nil {
			input[i], input[j] = input[j], nil
			for input[j] == nil {
				j--
			}
		}

		result += i * input[i].ID
	}

	for i := j; i < len(input); i++ {
		if input[i] == nil {
			continue
		}

		result += i * input[i].ID
	}

	return result
}

func part2(fp string) int {
	return 0
}

func printDisk(d []*File) {
	for _, file := range d {
		if file == nil {
			print(".")
		} else {
			print(file.ID)
		}
	}
	println()
}

func loadInput(fp string) []*File {
	nodes := make([]*File, 0)

	data, _ := os.ReadFile(fp)
	for i, id := 0, 0; i < len(data)-1; i, id = i+2, id+1 {
		fileSize := int(data[i] - '0')
		spaceSize := int(data[i+1] - '0')

		file := File{
			ID:   id,
			Size: fileSize,
		}

		for j := 0; j < fileSize; j++ {
			nodes = append(nodes, &file)
		}

		for j := 0; j < spaceSize; j++ {
			nodes = append(nodes, nil)
		}
	}

	return nodes
}
