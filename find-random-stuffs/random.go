package main

import (
	"bufio"
	"fmt"
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

func decodeLine(line string) []int {
	numbers := strings.Split(strings.TrimSpace(line), " ")

	n1, err := strconv.Atoi(numbers[0])
	check(err)

	n2, err := strconv.Atoi(numbers[1])
	check(err)

	var coll []int

	for i := n1; i <= n2; i++ {
		if i%7 == 0 && i%5 != 0 {
			coll = append(coll, i)
		}
	}
	return coll
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
		fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(decodeLine(scanner.Text())), " ", ",", -1), "[]"))
	}
}
