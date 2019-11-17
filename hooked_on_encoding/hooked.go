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

	// read it for scanning line by line
	scanner := bufio.NewScanner(file)
	var line string

	// string builder to accumulate output in
	var out strings.Builder

	// loop over lines until they exist
	for scanner.Scan() {
		// decode the current line
		line = strings.TrimSpace(scanner.Text())

		// slice holding each section delimited by spaces
		sections := strings.Split(line, " ")

		// range over sections
		for _, s := range sections {
			// this matches first word of each line
			if len(s) != 0 && !strings.ContainsAny(s, "-'") {
				out.WriteString(fmt.Sprintf("%s", string(s[0])))
				// the rest of characters are extracted from remaining sections
			} else {
				words := strings.Split(s, "-")
				for _, w := range words {
					// if no apostrophe, just split and write first char
					if len(w) != 0 && !strings.ContainsAny(w, "'") {
						out.WriteString(fmt.Sprintf("%s", string(w[0])))
						// if apostrophe then split and write first char of each half
					} else {
						spl := strings.Split(w, "'")
						if len(spl[0]) != 0 && len(spl[1]) != 0 {
							out.WriteString(fmt.Sprintf("%v'%v", string(spl[0][0]), string(spl[1][0])))
						}
					}
					// add in any commas if they existed in section
					if strings.ContainsAny(w, ",") {
						out.WriteString(",")
					}
				}
			}
			// add in spaces between words
			out.WriteString(" ")
		}
		// add in newlines between words
		out.WriteString("\n")
	}
	// quick and dirty way of cleaning the output from any extra spaces...
	cleaner := bufio.NewScanner(strings.NewReader(out.String()))
	for cleaner.Scan() {
		fmt.Printf("%s\n", strings.TrimSpace(cleaner.Text()))
	}
}
