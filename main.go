package main

import (
	"fmt"

	iotloop "constellationlabs.com/iotloop/iotnet"
)

func main() {
	fmt.Println(iotloop.Hello())
	c := make(chan bool)
	go iotloop.ListenForToken(2468, c)
	running := true
	for running {
		select {
		case rxok := <-c:
			if rxok {
				fmt.Println("token received")
			} else {
				fmt.Println("problem receiveing token")
			}
		}
	}
}
