package main

import (
	"fmt"
	"time"

	"constellationlabs.com/iotloop/iotnet"
	"constellationlabs.com/iotloop/ledcontrol"
)

const portnumber = 2468

func main() {
	fmt.Println(iotnet.Myname())
	if "pi-cluster-00" == iotnet.Myname() {
		// start the whole loop by sending an initial token
		iotnet.PassToken(iotnet.GetNextHop(), portnumber)
	}
	c := make(chan bool)
	go iotnet.ListenForToken(portnumber, c)
	timeout := make(chan bool, 1)
	go waittimeout(timeout)
	running := true
	for running {
		select {
		case rxok := <-c:
			if rxok {
				fmt.Println("token received, passing to " + iotnet.GetNextHop())
				newcolors := []uint32{0x440000, 0x004400, 0x000044, 0x222222} // using GGRRBB color ordering
				ledcontrol.SetLeds(newcolors)
				if iotnet.PassToken(iotnet.GetNextHop(), portnumber) {
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
	time.Sleep(30 * time.Second)
	timeoutchan <- true
}
