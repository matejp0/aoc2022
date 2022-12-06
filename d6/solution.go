package d6

import (

	"github.com/matejp0/aoc2022/utils"
)

var line = utils.Read()[0]

func Part1() int {
  return exclusiveIndex(line, 4)
}

func Part2() int {
  return exclusiveIndex(line, 14)
}

func exclusiveIndex(line string, count int) int {
  mapa := make(map[byte]uint)

  nextchar:
  for i := range line {
    mapa[line[i]]++

    if i >= count {
      mapa[line[i-count]]--

      for _, count := range mapa {
        if count > 1 {
          continue nextchar
        }
      }

      return i+1
    }
  }
  return 0
}
