package iotloop

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
