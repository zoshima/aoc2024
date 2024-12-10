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
	input := loadInput(fp)

	j := len(input) - 1
	for i := 0; i < j; i++ {
		if input[i] == nil {
			input[i], input[j] = input[j], nil
			for input[j] == nil {
				j--
			}
		}
	}

	return checksum(input)
}

func checksum(files []*File) int {
	result := 0
	for i := 0; i < len(files); i++ {
		if files[i] == nil {
			continue
		}

		result += i * files[i].ID
	}

	return result
}

func part2(fp string) int {
	input := loadInput(fp)

	for i := len(input) - 1; i > 0; i-- {
		file := input[i]
		if file == nil {
			continue
		}

		for j, slotSize := 0, 0; j < i; j++ {
			if input[j] != nil {
				slotSize = 0
				continue
			}

			slotSize++
			if slotSize == file.Size {
				for k := 0; k < slotSize; k++ {
					input[j-k], input[i-k] = input[i-k], input[j-k]
				}
				break
			}
		}
	}

	return checksum(input)
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
