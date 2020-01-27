package ledcontrol

import (
	"fmt"

	"github.com/jgarff/rpi_ws281x/golang/ws2811"
)

const (
	pin        = 18
	count      = 4
	brightness = 255
)

// SetLeds sets a string of RGB LEDs to an array of colors in u32 format
func SetLeds(colors []uint32) {
	defer ws2811.Fini()
	err := ws2811.Init(pin, count, brightness)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < len(colors); i++ {
		ws2811.SetLed(i, colors[i])
		err := ws2811.Render()
		if err != nil {
			ws2811.Clear()
			fmt.Println(err)
		}
	}
}
