package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// error handling function
func check(e error) {
	if e != nil {
		panic(e)
	}
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

	// run scanner on input
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			fmt.Println(scanner.Text())
		}
	}
}
