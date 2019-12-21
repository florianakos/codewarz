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
		var counter, sum int
		for _, e := range(strings.Split(strings.TrimSpace(scanner.Text()), " ")) {
			counter += 1
			n, _ := strconv.Atoi(e)
			sum += n
		}
		fmt.Printf("%d\n", counter + sum)
	}
}
