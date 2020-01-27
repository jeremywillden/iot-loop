package main

import (
	"fmt"

	"constellationlabs.com/iotloop/iotnet"
	"constellationlabs.com/iotloop/ledcontrol"
)

func main() {
	fmt.Println(iotnet.Hello())
	c := make(chan bool)
	go iotnet.ListenForToken(2468, c)
	running := true
	for running {
		select {
		case rxok := <-c:
			if rxok {
				fmt.Println("token received")
				newcolors := []uint32{0xFF0000, 0x00FF00}
				ledcontrol.SetLeds(newcolors)

			}
		}
	}
}
