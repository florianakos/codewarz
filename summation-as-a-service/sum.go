package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

// error handling function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func decodeLine(line string) string {
  numbers := strings.Split(strings.TrimSpace(line), " ")

	n1, err := strconv.Atoi(numbers[0])
	check(err)

	n2, err := strconv.Atoi(numbers[1])
	check(err)

	var sum int

	if n1 < n2 {
		for i := n1; i <= n2; i++ {
			// fmt.Printf("%d ", i)
			sum += i
		}
	} else {
		for i := n2; i <= n1; i++ {
			// fmt.Printf("%d ", i)
			sum += i
		}
	}
	return fmt.Sprintf("%d", sum)
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

	scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    fmt.Println(decodeLine(scanner.Text()))
  }
}
