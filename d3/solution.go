package d3

import (
	"github.com/matejp0/aoc2022/utils"
)

var input []string = utils.Read()

func Part1() int {
	var sum int

	for _, line := range input {
		out:
		for _, char := range line[:len(line)/2] {
			for _, char2 := range line[len(line)/2:] {
				if char == char2 {
					sum += getScore(char)
					break out
				}
			}
		}
	}

	return sum
}

func Part2() int {
	var sum int

	for i := 0; i < len(input)-2; i+=3 {
		var temp string
		for _, char1 := range input[i] {
			for _, char2 := range input[i+1] {
				if char1 == char2 {
					temp += string(char1)
				}
			}
		}

		out:
		for _, char := range temp {
			for _, char2 := range input[i+2] {
				if char == char2 {
					sum += getScore(char)
					break out
				}
			}
		}
	}

	return sum
}

func getScore(char rune) int {
	if c := char - 'A'; c <= 26 { // it's uppercase
		return int(c) +1 + 26
	} else {
		return int(char - 'a') + 1
	}
}
