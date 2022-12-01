package utils

import (
	"bufio"
	"log"
	"os"
	"runtime"
	"strings"
)

func Read() []string {
  // Get the name of the package
  pc, _, _, _ := runtime.Caller(1)
  funcName := runtime.FuncForPC(pc).Name()
	lastSlash := strings.LastIndexByte(funcName, '/')
	if lastSlash < 0 {
		lastSlash = 0
	}
	lastDot := strings.LastIndexByte(funcName[lastSlash:], '.') + lastSlash


  file, err := os.Open(funcName[lastDot-2:lastDot] + "/input")
  if err != nil { log.Fatal(err) }
  defer file.Close()

  lines := make([]string, 0)

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  if err := scanner.Err(); err != nil { log.Fatal(err) }

  return lines
}
