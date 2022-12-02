package d2

import (
	"strings"

	"github.com/matejp0/aoc2022/utils"
)

var input []string = utils.Read()

func Part1() int {
	var score int
	for _, line := range input {
		arr := strings.Fields(line)
		opponent := int(arr[0][0])-'A'
		me := int(arr[1][0])-'X'

		score += me + 1

		if modLikePython(me - opponent, 3) == 1 {
			score += 6
		} else if me == opponent {
			score += 3
		}
	}

	return score
}

func modLikePython(d, m int) int {
   var res int = d % m
   if (res < 0 && m > 0) {
      return res + m
   }
   return res
}

func Part2() int {
	var score int
	for _, line := range input {
		arr := strings.Fields(line)

		switch arr[1][0] {
		case 'X':
			score += modLikePython(int(arr[0][0])-'A'-1, 3)+1
		case 'Y':
			score += int(arr[0][0])-'A'+1+3
		case 'Z':
			score += modLikePython(int(arr[0][0])-'A'+1, 3)+1+6
		}

	}

	return score
}
