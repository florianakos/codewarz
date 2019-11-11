package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"
)

// global dict holding the contents of the dictionary file
var wordDict = make(map[string]bool)

// error handling function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// func that loads the dict into memory
func loadDictionary() {
	var dictLoc string
	// if running on MacOS then use local dict in repository folder
	if runtime.GOOS == "darwin" {
		dictLoc = "american-english"
		// else use the Unix default dict
	} else {
		dictLoc = "/usr/share/dict/american-english"
	}
	// open the determined file location
	dict, err := os.Open(dictLoc)
	check(err)
	defer dict.Close()
	// read it via bufio scanner
	ds := bufio.NewScanner(dict)
	for ds.Scan() {
		// import lines of dict into string map
		wordDict[ds.Text()] = true
	}
}

// new version of dict lookup via direct access to string map storing each line
func dictLookupFast(word string) bool {
	if _, ok := wordDict[word]; ok {
		return true
	}
	return false
}

// helper func to look up a given string in dictionary (location depends on OS)
func dictLookup(word string) bool {
	var dictLoc string
	if runtime.GOOS == "darwin" {
		dictLoc = "american-english"
	} else {
		dictLoc = "/usr/share/dict/american-english"
	}
	_, err := exec.Command("grep", "-w", word, dictLoc).Output()
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

func workerFunc(jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	// Decreasing internal counter for wait-group as soon as goroutine finishes
	defer wg.Done()

	// eventually I want to have a []string channel to work on a chunk of lines not just one line of text
	for j := range jobs {
		start := time.Now()
		line := decode(j)

		// test and further decode until valid work is found
		for !(dictLookupFast(line)) {
			line = decode(line)
		}
		elapsed := time.Since(start)
		results <- fmt.Sprintf("%s \t%s", line, elapsed)
	}
}

func main() {
	// check if filename in arg is present
	if (len(os.Args) - 1) != 1 {
		fmt.Println("ERROR: missing argument!")
		os.Exit(1)
	}

	// load OS specific dictionary into memory for fast execution
	loadDictionary()

	// open file for reading
	f, err := os.Open(os.Args[1:][0])
	check(err)

	input, err := ioutil.ReadAll(f)
	check(err)

	file := strings.NewReader(string(input))

	// do I need buffered channels here?
	jobs := make(chan string)
	results := make(chan string)

	// global waitgroup
	wg := new(sync.WaitGroup)

	// start up some workers that will block and wait
	for w := 1; w <= 5; w++ {
		wg.Add(1)
		go workerFunc(jobs, results, wg)
	}

	// Go over a file line by line and queue them up
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			jobs <- scanner.Text()
		}
		close(jobs)
	}()

	// Now collect all the results...
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print out the results from the results channel.
	for v := range results {
		fmt.Println(v)
	}
}
