package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	var err error

	// Create an ls command.
	ls := exec.Command("more", "/usr/share/dict/american-english")

	// Create a grep command that searches for anything
	// that contains .go in it's filename.
	grep := exec.Command("grep", "interpretationss")

	// Set grep's stdin to the output of the ls command.
	grep.Stdin, err = ls.StdoutPipe()
	if err != nil {
		log.Fatalln(err)
	}

	// Set grep's stdout to os.Stdout
	grep.Stdout = os.Stdout

	// Start the grep command first. (The order will be last command first)
	must(grep.Start())

	// Run the ls command. (Run calls start and also calls wait)
	must(ls.Run())

	// Wait for the grep command to finish.
	err = grep.Wait()
  if err != nil {
    log.Println("Not found")
  }
}

// Must checks if an error is nil and if it isn't calls log.Fatalln(err)
func must(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
