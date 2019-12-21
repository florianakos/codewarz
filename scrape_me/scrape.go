package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
	res, err := http.Get(os.Args[1])
	check(err)
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	check(err)
	fmt.Printf("%s", body)
}
