package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
  "encoding/hex"
)

// error handling function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convertHex(hx string) string {
    bs, err := hex.DecodeString(strings.Split(hx, "x")[1])
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
    fmt.Println(err)
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

  // run scanner on input
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
    tokens := strings.Split(strings.TrimSpace(scanner.Text()), " ")
    // parse each token as either hex, binary, octal or decial
    for _, val := range tokens {
      if len(val) != 0 {
        // token is hexadecimal string
        if strings.Contains(val, "x") {
          fmt.Print(convertHex(val))
        // token is binary string
        } else if strings.Contains(val, "b") {
          fmt.Print(convertBinary(val))
        // token is octal string
        } else if strings.Index(val, "0") == 0 {
          fmt.Print(convertOctal(val))
        // token is decimal string
        } else {
          tmp, _ :=  strconv.Atoi(val)
          fmt.Print(string(tmp))
        }
      }
    }
	}
}
