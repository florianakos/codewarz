package main

import (
	"bufio"
	"encoding/hex"
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

func convertHex(hx string) string {
	str := strings.Split(hx, "x")[1]
	// deal with error: encoding/hex: odd length hex string
	if len(str) == 1 {
		str = "0" + str
	}
	bs, err := hex.DecodeString(str)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(bs)
}

func convertBinary(bin string) string {
	input := strings.Split(bin, "b")[1]
	b := make([]byte, 0)
	for _, s := range strings.Fields(input) {
		n, _ := strconv.ParseUint(s, 2, 8)
		b = append(b, byte(n))
	}
	return string(b)
}

func convertOctal(oct string) string {
	output, err := strconv.ParseInt(oct, 8, 64)
	if err != nil {
		// fmt.Println(err)
		return ""
	}
	return string(output)
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

	//
	var out strings.Builder

	// run scanner on input
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tokens := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		// parse each token as either hex, binary, octal or decial
		for _, v := range tokens {
			if len(v) != 0 {
				val := strings.TrimSpace(v)
				// token is hexadecimal string
				if strings.Contains(val, "x") {
					out.WriteString(fmt.Sprint(convertHex(val)))
					// token is binary string
				} else if strings.Contains(val, "b") {
					out.WriteString(fmt.Sprint(convertBinary(val)))
					// token is octal string
				} else if strings.Index(val, "0") == 0 {
					out.WriteString(fmt.Sprint(convertOctal(val)))
					// token is decimal string
				} else {
					tmp, _ := strconv.Atoi(val)
					out.WriteString(fmt.Sprint(string(tmp)))
				}
			}
		}
	}
	fmt.Print(out.String())
}
