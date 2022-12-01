package d1

import (
	"log"
	"sort"
	"strconv"

	"github.com/matejp0/aoc2022/utils"
)

var input []string = utils.Read()

func Part1() int {
  var temp int
  var max int
  for _, line := range input {
    if line == "" { 
      if temp > max  { max = temp }
      temp = 0
      continue
    }
    n, err := strconv.Atoi(line)
    if err != nil {
      log.Fatal(err)
    }
    temp += n
  }
  if temp > max  { max = temp }

  return max
}

func Part2() int {
  list := make([]int, 0)
  var temp int
  for _, line := range input {
    if line == "" {
      list = append(list, temp)
      temp = 0
      continue
    }

    n, err := strconv.Atoi(line)
    if err != nil {
      log.Fatal(err)
    }
    temp += n
  }
  list = append(list, temp)

  sort.Slice(list, func(i, j int) bool {
    return list[i]>list[j]
  })

  return list[0]+list[1]+list[2]
}
