package main

import "os"
import "fmt"

func main() {

	argsWithoutProg := os.Args[1:]

	for _, i := range argsWithoutProg {
		fmt.Printf("%s ", i)
	}
}
