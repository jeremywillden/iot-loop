package iotnet

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

// Myname is a simple Hello World test function
func Myname() string {
	name, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return name
}

// GetNextHop looks up this device hostname and uses it to find the IP address of the next device
func GetNextHop() string {
	nexthops := map[string]string{
		"pi-cluster-00": "10.0.1.61", // pi-cluster-00 is 10.0.1.60, never receives
		"pi-cluster-01": "10.0.1.62",
		"pi-cluster-02": "10.0.1.63",
		"pi-cluster-03": "10.0.1.64",
		"pi-cluster-04": "10.0.1.65",
		"pi-cluster-05": "10.0.1.66",
		"pi-cluster-06": "10.0.1.67",
		"pi-cluster-07": "10.0.1.68",
		"pi-cluster-08": "10.0.1.69",
		"pi-cluster-09": "10.0.1.70",
		"pi-cluster-10": "10.0.1.71",
		"pi-cluster-11": "10.0.1.72",
		"pi-cluster-12": "10.0.1.73",
		"pi-cluster-13": "10.0.1.74",
		"pi-cluster-14": "10.0.1.75",
		"pi-cluster-15": "10.0.1.76",
		"pi-cluster-16": "10.0.1.77",
		"pi-cluster-17": "10.0.1.78",
		"pi-cluster-18": "10.0.1.79",
		"pi-cluster-19": "10.0.1.80",
		"pi-cluster-20": "10.0.1.81",
		"pi-cluster-21": "10.0.1.82",
		"pi-cluster-22": "10.0.1.83",
		"pi-cluster-23": "10.0.1.84",
		"pi-cluster-24": "10.0.1.61",
	}
	/*nexthops := map[string]string{
		"pi-cluster-00": "10.100.205.101", // pi-cluster-00 is 10.100.205.100, never receives
		"pi-cluster-01": "10.100.205.102",
		"pi-cluster-02": "10.100.205.103",
		"pi-cluster-03": "10.100.205.104",
		"pi-cluster-04": "10.100.205.105",
		"pi-cluster-05": "10.100.205.106",
		"pi-cluster-06": "10.100.205.107",
		"pi-cluster-07": "10.100.205.108",
		"pi-cluster-08": "10.100.205.109",
		"pi-cluster-09": "10.100.205.110",
		"pi-cluster-10": "10.100.205.111",
		"pi-cluster-11": "10.100.205.112",
		"pi-cluster-12": "10.100.205.113",
		"pi-cluster-13": "10.100.205.114",
		"pi-cluster-14": "10.100.205.115",
		"pi-cluster-15": "10.100.205.116",
		"pi-cluster-16": "10.100.205.117",
		"pi-cluster-17": "10.100.205.118",
		"pi-cluster-18": "10.100.205.119",
		"pi-cluster-19": "10.100.205.120",
		"pi-cluster-20": "10.100.205.121",
		"pi-cluster-21": "10.100.205.122",
		"pi-cluster-22": "10.100.205.123",
		"pi-cluster-23": "10.100.205.124",
		"pi-cluster-24": "10.100.205.101",
	}*/
	return nexthops[Myname()]
}

// PassToken sends the token on to the next device in the chain, returns true if successful
func PassToken(passto string, portnum int) (passed bool) {
	portstr := strconv.Itoa(portnum)
	conn, err := net.Dial("tcp", passto+":"+portstr)
	if err == nil {
		fmt.Fprintf(conn, "TOKEN\r\n")
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err == nil {
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
