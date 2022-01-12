package Adapter


type Television interface {
	volumeUp() int
	volumeDown() int
	turnOn()
	setChannel(ch int)
}
