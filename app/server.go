package main

import (
	"fmt"
	"net"
	"os"
)

func readDataFromConn(con_p *net.Conn) string {
	data := make([]byte, 0)
	con := *con_p
	_, err := con.Read(data)
	if err != nil {
		panic(err.Error())
	}
	return string(data[:])
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	con, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	fmt.Println(readDataFromConn(&con))
	resp := "+PONG\r\n"
	con.Write([]byte(resp))
	con.Close()
}
