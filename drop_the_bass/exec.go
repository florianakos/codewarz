package main

import "os/exec"
import "fmt"

func testWord(word string) bool {
	_, err := exec.Command("grep", "-w", word, "/usr/share/dict/american-english").Output()
	if err != nil {
		return false
	}
	return true
}

func main() {
	fmt.Println(testWord("interprett"))
}
