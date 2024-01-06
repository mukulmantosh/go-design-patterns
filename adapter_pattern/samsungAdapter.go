package main

type samsungAdapter struct {
	stv *SamsungTV
}

func (ss *samsungAdapter) turnOn() {
	ss.stv.setOnState(true)
}

func (ss *samsungAdapter) turnOff() {
	ss.stv.setOnState(false)
}

func (ss *samsungAdapter) volumeUp() int {
	vol := ss.stv.getVolume() + 1
	ss.stv.setVolume(vol)
	return vol
}

func (ss *samsungAdapter) volumeDown() int {
	vol := ss.stv.getVolume() - 1
	ss.stv.setVolume(vol)
	return vol
}

func (ss *samsungAdapter) channelUp() int {
	ch := ss.stv.getChannel() + 1
	ss.stv.setChannel(ch)
	return ch
}

func (ss *samsungAdapter) channelDown() int {
	ch := ss.stv.getChannel() - 1
	ss.stv.setChannel(ch)
	return ch
}

func (ss *samsungAdapter) goToChannel(ch int) {
	ss.stv.setChannel(ch)
}
