package main

import (
	"bufio"
	// "encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// simple error handling function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// func to find max in slice of ints
func max(slc []int) int {
  base := 0
  for _, v := range slc {
    if v > base {
      base = v
    }
  }
  return base
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

	// var out strings.Builder
  var inputNumbers []int

	// read lines into int slice for later processing
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
    num, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
    inputNumbers = append(inputNumbers, num)
	}

  // increment counter slice element for every group of 10
  counterSlice := make([]int, (max(inputNumbers)/10+1))
  for _, val := range inputNumbers {
    counterSlice[ val / 10 ] += 1
  }

  // print out the slice as requested
  for _, v := range counterSlice {
    fmt.Printf("%d", v)
  }
}
