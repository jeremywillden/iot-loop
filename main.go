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
				newcolors := []uint32{0x220000, 0x002200} // using GGRRBB color ordering
				ledcontrol.SetLeds(newcolors)
				if iotnet.PassToken("10.0.1.60", 2468) {
					fmt.Println("token passed")
					offcolors := []uint32{0x000000, 0x000000} // using GGRRBB color ordering
					ledcontrol.SetLeds(offcolors)
				}
			}
		}
	}
}
