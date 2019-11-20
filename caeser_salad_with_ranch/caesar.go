package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

func caesar(r rune, shift int) rune {
	s := int(r) + shift
	// Special cases for ignoring punctuation.
	if unicode.IsPunct(r) {
		return rune(r)
		// If Capital Letter
	} else if unicode.IsUpper(r) {
		if s > 'Z' {
			return rune(s - 26)
		} else if s < 'A' {
			return rune(s + 26)
		}
		// if Lower Case letter
	} else {
		if s > 'z' {
			return rune(s - 26)
		} else if s < 'a' {
			return rune(s + 26)
		}
	}
	return rune(s)
}

func main() {
	// check if filename in arg is present
	if (len(os.Args) - 1) != 1 {
		fmt.Println("ERROR: missing argument!")
		os.Exit(1)
	}

	// read whole file as text into byte array
	byteString, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Print(err)
	}

	// clean and split input to individual lines
	lines := strings.Split(strings.TrimSpace(string(byteString)), "\n")

	// start offset at -1 coz first word is ignored
	offset := -1
	for i, l := range lines {
		// split line into words and print first word as is
		words := strings.Split(l, " ")
		fmt.Printf("%s", words[0])

		// iterate over the remaining words and process accordingly
		for j, w := range words[1:] {
			// if not last word then add space
			if j < len(words) {
				fmt.Printf(" %v", strings.Map(func(r rune) rune {
					return caesar(r, offset)
				}, w))
				// if last word then dont add space!!!!
			} else {
				fmt.Printf("%v", strings.Map(func(r rune) rune {
					return caesar(r, offset)
				}, w))
			}
			// decrease word offset for the caesar cipher
			offset--
		}
		// reset offset for a new line
		offset--
		if i != len(lines)-1 {
			fmt.Println()
		}
	}
}
