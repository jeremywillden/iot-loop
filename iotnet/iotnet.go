package iotnet

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

// Hello is a simple Hello World test function
func Hello() string {
	return "hello world!"
}

// PassToken sends the token on to the next device in the chain, returns true if successful
func PassToken(passto string, portnum int) (passed bool) {
	portstr := strconv.Itoa(portnum)
	conn, err := net.Dial("tcp", passto+":"+portstr)
	if err == nil {
		fmt.Fprintf(conn, "TOKEN\r\n")
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if strings.Contains(response, "OK") {
				return true
			}
		}
	}
	return false
}

// ListenForToken opens a connection, returning a boolean true in a channel when the token is received
func ListenForToken(portnum int, tokenchan chan bool) {
	running := true
	portstr := strconv.Itoa(portnum)
	rxsocket, err := net.Listen("tcp", ":"+portstr)
	if err != nil {
		running = false
	}
	for running {
		conn, err := rxsocket.Accept()
		defer conn.Close()
		if err != nil {
			running = false
		}
		go receiveToken(conn, tokenchan)
	}
}

func receiveToken(conn net.Conn, tokenchan chan bool) {
	buffer := make([]byte, 1024)
	len, err := conn.Read(buffer)
	if (err == nil) && (len < 128) {
		conn.Write([]byte("OK\r\n"))
		if strings.Contains(string(buffer), "TOKEN") {
			tokenchan <- true
		}
	}
	tokenchan <- false
	conn.Close()
}
