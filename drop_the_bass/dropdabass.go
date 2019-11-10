package main

import "fmt"
import "os"
//import "io/ioutil"
import "bufio"
import "encoding/base64"
import "log"
import "os/exec"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func decode(src string) string {
  strToDecode := src
  var data []byte
	//fmt.Printf("\ndecoding: %s\n", strToDecode)
  data, err := base64.StdEncoding.DecodeString(strToDecode)
  for {
    if err != nil {
  		//fmt.Println("error:", err)
  		strToDecode += "="
  	} else {
      break
    }
    data, err = base64.StdEncoding.DecodeString(strToDecode)
  }

	//fmt.Printf("decoded:  %s\n\n", string(data))
	//dst :=  fmt.Sprintf("%s\n", data)
	return string(data)
}

func search(toSearch string) bool {
	var err error

	// Create an ls command.
	ls := exec.Command("more", "/usr/share/dict/american-english")

	// Create a grep command that searches for anything
	// that contains .go in it's filename.
	grep := exec.Command("grep", toSearch)

	// Set grep's stdin to the output of the ls command.
	grep.Stdin, err = ls.StdoutPipe()
	if err != nil {
		log.Fatalln(err)
	}

	// Start the grep command first. (The order will be last command first)
	must(grep.Start())

	// Run the ls command. (Run calls start and also calls wait)
	must(ls.Run())

	// Wait for the grep command to finish.
  err = grep.Wait()
  if err != nil {
    return false
  }
  return true
}

// Must checks if an error is nil and if it isn't calls log.Fatalln(err)
func must(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
  //line := ""
  var d string
  //line1 := ""
  if (len(os.Args)-1) == 1 {
	    file, err := os.Open(os.Args[1:][0])
	    check(err)

    	scanner := bufio.NewScanner(file)

      for scanner.Scan() {
    		line := scanner.Text()
        d = decode(line)
        verdict := search(d)
        for verdict == false {
          d = decode(d)
          verdict := search(d)
          if verdict == true {
            break
          }
        }
        fmt.Println(d)
      }



    	//line = decode(line)
    	//line = decode(line)
    	//line = decode(line)
    	//line = decode(line)
    	//line = decode(line)
    	//verdict := search(decode(line))

    	//fmt.Printf("\nfinal: %v\n",verdict)

  } else {
      fmt.Println("ERROR")
  }
}
