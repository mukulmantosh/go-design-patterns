package main

import "fmt"

func main() {
	// Create instance of the two TV types with some default values

	tv1 := &SamsungTV{
		currentChan:   13,
		currentVolume: 35,
		tvOn:          true,
	}

	tv2 := &SonyTV{
		vol:     20,
		channel: 9,
		isOn:    true,
	}

	// Because the SonyTV implements the "television" interface
	tv2.turnOn()
	tv2.volumeUp()
	tv2.volumeDown()
	tv2.channelUp()
	tv2.channelDown()
	tv2.goToChannel(68)
	tv2.turnOff()

	fmt.Println("=*=*=*=*=*=*=*=*=*")

	// Creating SamsungTV Adapter
	samAdapter := &samsungAdapter{stv: tv1}
	samAdapter.turnOn()
	samAdapter.volumeUp()
	samAdapter.volumeDown()
	samAdapter.channelUp()
	samAdapter.channelDown()
	samAdapter.goToChannel(56)
	samAdapter.turnOff()
}
