package main

import (
	"bufio"
	"encoding/hex"
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

// funciton that turns hex string into normal string (1 character)
func decodeChar(ch string) string {
	decoded, err := hex.DecodeString(ch)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s", decoded)
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
		chars := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		for _, val := range chars {
			fmt.Print(decodeChar(val))
		}
	}
}
