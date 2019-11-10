package main

import "fmt"
import "os"
import "os/exec"
import "bufio"
import "encoding/base64"
import "runtime"
import "sync"

// error handling function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// helper func to look up a given string in dictionary (location depends on OS)
func dictLookup(word string) bool {
	var dictLocation string
	if runtime.GOOS == "darwin" {
    dictLocation = "american-english"
	} else {
		dictLocation = "/usr/share/dict/american-english"
	}
	_, err := exec.Command("grep", "-w", word, dictLocation).Output()
	if err != nil {
		return false
	}
	return true
}

// helper func to decode a base64 encoded string (append necessary = at end to make it valid...)
func decode(src string) string {
	data, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		src2 := src + "="
		data2, err2 := base64.StdEncoding.DecodeString(src2)
		if err2 != nil {
			src3 := src2 + "="
			data3, _ := base64.StdEncoding.DecodeString(src3)
			return string(data3)
		}
		return string(data2)
	}
	return string(data)
}

func main() {
	// global variable to enable concurrency
	var wg sync.WaitGroup

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

	// loop over lines until they exist
	for scanner.Scan() {

		wg.Add(1)
		go func() {
			// decode the current line
			line = decode(scanner.Text())

			// test and further decode until valid work is found
			for !(dictLookup(line)) {
				line = decode(line)
			}
			fmt.Println(line)
      wg.Done()
		}()
		wg.Wait()
	}
}
