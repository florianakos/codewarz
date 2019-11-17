package main

import (
	"bufio"
	_"encoding/hex"
	"fmt"
	"os"
	_"strings"
)

// error handling function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// function that reverses a string
func backwards(input string) string {
  var tempStr string
  for i := len(input)-1; i > -1; i-- {
    tempStr += string(input[i])
  }
  return tempStr
}

func main() {
	// check if filename in arg is present
	if (len(os.Args) - 1) != 1 {
		fmt.Println("ERROR: missing argument!")
		os.Exit(1)
	}

	// open file for reading
	file, err := os.Open(os.Args[1:][0])
	check(err)

	// read it for scanning line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
    forwardTxt := scanner.Text()
    if forwardTxt == backwards(forwardTxt) {
      fmt.Println("True")
    } else {
      fmt.Println("False")
    }
	}
}
