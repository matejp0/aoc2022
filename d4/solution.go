package d4

import (
	"log"
	"strconv"
	"strings"

	"github.com/matejp0/aoc2022/utils"
)

var input = parse(utils.Read())


func Part1() int {
	var contained int

	for i := 0; i < len(input)-1; i+=2 {
		elf1 := input[i]
		elf2 := input[i+1]

		if (elf1[0] <= elf2[0] && elf1[1] >= elf2[1]) || (elf2[0] <= elf1[0] && elf2[1] >= elf1[1]) {
			contained++

		}
	}

	return contained
}

func Part2() int {
	var overlap int

	for i := 0; i < len(input)-1; i+=2 {
		elf1 := input[i]
		elf2 := input[i+1]

		out:
		for i := elf1[0]; i <= elf1[1]; i++ {
			for j := elf2[0]; j <= elf2[1]; j++ {
				if i == j {
					overlap++
					break out
				}

			}
		}

	}

	return overlap
}


func parse(input []string) [][2]int {
	result := make([][2]int, 0)

	for _, line := range input {
		for _, elf := range strings.Split(line, ",") {
			arr := strings.Split(elf, "-")
			int1, err := strconv.Atoi(arr[0])
			if err != nil {
				log.Fatal(err)
			}
			int2, err := strconv.Atoi(arr[1])
			if err != nil {
				log.Fatal(err)
			}
			result = append(result, [2]int{int1, int2})
		}
	}

	return result
}
