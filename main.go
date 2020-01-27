package main

import (
	"fmt"

	"constellationlabs.com/iotloop/iotnet"
	"constellationlabs.com/iotloop/ledcontrol"
)

func main() {
	fmt.Println(iotnet.Myname())
	c := make(chan bool)
	go iotnet.ListenForToken(2468, c)
	running := true
	for running {
		select {
		case rxok := <-c:
			if rxok {
				fmt.Println("token received, passing to " + iotnet.GetNextHop())
				newcolors := []uint32{0x110000, 0x001100} // using GGRRBB color ordering
				ledcontrol.SetLeds(newcolors)
				if iotnet.PassToken(iotnet.GetNextHop(), 2468) {
					fmt.Println("token passed")
					offcolors := []uint32{0x000000, 0x000000} // using GGRRBB color ordering
					ledcontrol.SetLeds(offcolors)
				}
			}
		}
	}
}
