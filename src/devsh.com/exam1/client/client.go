package main

import (
	"net"
	"fmt"
	"os"
	"bufio"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:9999")
	// read in input from stdin
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	// send to socket
	fmt.Fprintf(conn, text + "\n")

	// listen for reply
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message from server:" + message)

	conn.Close()
}