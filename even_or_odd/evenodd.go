package main

import (
	"bufio"
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

func main() {
	if (len(os.Args) - 1) != 1 {
		fmt.Println("ERROR: missing argument!")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1:][0])
	check(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if n%2 == 0 {
			fmt.Printf("%d True\n", n)
		} else {
			fmt.Printf("%d False\n", n)
		}
	}
}
