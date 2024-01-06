package main

import "fmt"

type SamsungTV struct {
	currentChan   int
	currentVolume int
	tvOn          bool
}

func (tv *SamsungTV) getVolume() int {
	fmt.Println("SamsungTV volume is", tv.currentVolume)
	return tv.currentVolume
}

func (tv *SamsungTV) setVolume(vol int) {
	fmt.Println("Setting SamsungTV volume to", vol)
	tv.currentVolume = vol
}

func (tv *SamsungTV) getChannel() int {
	fmt.Println("SamsungTV channel is ", tv.currentChan)
	return tv.currentChan
}

func (tv *SamsungTV) setChannel(ch int) {
	fmt.Println("Setting SamsungTV channel to", ch)
	tv.currentChan = ch
}

func (tv *SamsungTV) setOnState(tvOn bool) {
	if tvOn == true {
		fmt.Println("SamsungTV is on")
	} else {
		fmt.Println("SamsungTV is off")
	}
	tv.tvOn = tvOn
}
