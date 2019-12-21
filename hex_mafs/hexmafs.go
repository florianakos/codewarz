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
		nums := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		var sum int64
		for _, val := range nums {
      i, err := strconv.ParseInt(strings.Split(val, "x")[1], 16, 64);
			if  err == nil {
				sum += i
			}
		}
		fmt.Println(sum)
	}
}
