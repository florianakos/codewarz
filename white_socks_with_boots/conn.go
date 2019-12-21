package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {
	if (len(os.Args) - 1) != 2 {
		fmt.Println("ERROR: missing 2 arguments!")
		os.Exit(1)
	}

	conn, _ := net.Dial("tcp", os.Args[1:][0]+":"+os.Args[1:][1])
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print(message)
}
