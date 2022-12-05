package d5

import (
	"log"
	"strconv"
	"strings"

	"github.com/matejp0/aoc2022/utils"
)


func Part1() string {
  var str string

  input, instructions := parse(utils.Read())

  for _, line := range instructions {
    arr := strings.Fields(line)

    count, err := strconv.Atoi(arr[1])
    if err != nil { log.Fatal(err) }
    
    from, err := strconv.Atoi(arr[3])
    if err != nil { log.Fatal(err) }

    to, err := strconv.Atoi(arr[5])
    if err != nil { log.Fatal(err) }

    for i := 0; i < count; i++ {
      input[to-1] = append(input[to-1], input[from-1][len(input[from-1])-1])
      input[from-1] = input[from-1][:len(input[from-1])-1]
    }

  }

  for _, stack := range input {
    if stack != nil {
      str += string(stack[len(stack)-1])
    }
  }

  return str
}

func Part2() string {
  var str string

  input, instructions := parse(utils.Read())

  for _, line := range instructions {
    arr := strings.Fields(line)

    count, err := strconv.Atoi(arr[1])
    if err != nil { log.Fatal(err) }
    
    from, err := strconv.Atoi(arr[3])
    if err != nil { log.Fatal(err) }

    to, err := strconv.Atoi(arr[5])
    if err != nil { log.Fatal(err) }

    input[to-1] = append(input[to-1], input[from-1][(len(input[from-1])-count):(len(input[from-1]))]...)
    input[from-1] = input[from-1][:len(input[from-1])-count]

  }

  for _, stack := range input {
    if stack != nil {
      str += string(stack[len(stack)-1])
    }
  }

  return str
}

func parse(input []string) ([][]byte, []string) {
  stacks := make([][]byte, (len(input[0])+1)/4)
  for i, line := range input {
    if !strings.ContainsRune(line, '[') {
      for _, stack := range stacks {
        for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
            stack[i], stack[j] = stack[j], stack[i]
        }
      }
      return stacks, input[i+2:] // skip the number line and the empty line
    }

    for x := 1; x < len(line); x+=4 {
      if char := line[x]; char != ' ' {
        stacks[(x-1)/4] = append(stacks[(x-1)/4], char)
      }
    }
  }

  return nil, nil
}
