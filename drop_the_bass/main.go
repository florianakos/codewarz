package main

import "fmt"
import "os"
import "os/exec"
import "bufio"
import "encoding/base64"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func dictLookup(word string) bool {
	_, err := exec.Command("grep", "-w", word, "/usr/share/dict/american-english").Output()
	if err != nil {
		return false
	}
	return true
}

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
	line := ""
	if (len(os.Args) - 1) == 1 {
		file, err := os.Open(os.Args[1:][0])
		check(err)
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line = decode(scanner.Text())
			for !(dictLookup(line)) {
				line = decode(line)
			}
			fmt.Println(line)
		}
	} else {
		fmt.Println("ERROR: missing argument!")
	}
}
