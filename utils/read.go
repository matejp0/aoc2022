package utils

import (
	"io/ioutil"
	"log"
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


  contents, err := ioutil.ReadFile(funcName[lastDot-2:lastDot] + "/input")
  if err != nil {
    log.Fatal(err)
    return nil
  }

  return strings.Split(string(contents), "\n")
}
