package main

import "fmt"

type SonyTV struct {
	vol     int
	channel int
	isOn    bool
}

func (st *SonyTV) turnOn() {
	fmt.Println("SonyTV is now on")
	st.isOn = true
}

func (st *SonyTV) turnOff() {
	fmt.Println("SonyTV is now off")
	st.isOn = false
}

func (st *SonyTV) volumeUp() int {
	st.vol++
	fmt.Println("Increasing SonyTV volume to", st.vol)
	return st.vol
}

func (st *SonyTV) volumeDown() int {
	st.vol--
	fmt.Println("Decreasing SonyTV volume to", st.vol)
	return st.vol
}

func (st *SonyTV) channelUp() int {
	st.channel++
	fmt.Println("Increasing SonyTV channel to", st.channel)
	return st.channel
}

func (st *SonyTV) channelDown() int {
	st.channel--
	fmt.Println("Decreasing SonyTV channel to", st.channel)
	return st.channel
}

func (st *SonyTV) goToChannel(ch int) {
	st.channel = ch
	fmt.Println("going to channel number ", st.channel)
}
