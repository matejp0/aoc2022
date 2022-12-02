package d2

import (
	"strings"

	"github.com/matejp0/aoc2022/utils"
)

var input []string = utils.Read()
// Rock, Paper, Scissors
var win = map[byte]byte{
	'A': 'Y',
	'B': 'Z',
	'C': 'X',
}

var lose = map[byte]byte{
	'A': 'Z',
	'B': 'X',
	'C': 'Y',
}

func Part1() int {
	var score int
	for _, line := range input {
		arr := strings.Fields(line)

		score += int(arr[1][0]) - 'X' + 1 // Add 1/2/3 depending on the letter

		if win[arr[0][0]] == arr[1][0] {
			score += 6
		} else if arr[1][0]-'X' == arr[0][0]-'A' {
			score += 3
		}

	}

	return score

}

func Part2() int {
	var score int
	for _, line := range input {
		arr := strings.Fields(line)

		switch arr[1][0] {
		case 'X':
			score += int(lose[arr[0][0]])-'X'+1
		case 'Y':
			score += int(arr[0][0])-'A'+1+3
		case 'Z':
			score += int(win[arr[0][0]])-'X'+1+6
		}

	}

	return score
}
