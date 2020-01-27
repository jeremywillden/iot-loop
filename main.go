package main

import (
	"fmt"
	"time"

	"constellationlabs.com/iotloop/iotnet"
	"constellationlabs.com/iotloop/ledcontrol"
)

func main() {
	fmt.Println(iotnet.Myname())
	c := make(chan bool)
	go iotnet.ListenForToken(2468, c)
	timeout := make(chan bool, 1)
	go waittimeout(timeout)
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
		case <-timeout:
			fmt.Println("timeout complete, exiting")
			running = false
		}
	}
}

func waittimeout(timeoutchan chan bool) {
	time.Sleep(10 * time.Second)
	timeoutchan <- true
}
