package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "strings"
)

func main() {
  // check if filename in arg is present
	if (len(os.Args) - 1) != 1 {
		fmt.Println("ERROR: missing argument!")
		os.Exit(1)
	}

  // read whole file as text into byte array
  byteString, err := ioutil.ReadFile(os.Args[1])
  if err != nil {
      fmt.Print(err)
  }

  // convert byte array to string
  strText := string(byteString)

  // extract sentences by splitting on dot
  sentences := strings.Split(strText, ".")

  // extract first word of each sentence, plus clean from whitespace.
  for _, v := range(sentences) {
    word := strings.Split(strings.TrimSpace(v), " ")[0]
    if word != "" {
      fmt.Printf("%s ", word)
    }
  }
}
