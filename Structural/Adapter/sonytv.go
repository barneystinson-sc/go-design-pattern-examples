package Adapter


type SonyTv struct{
	currentVol int
	currentChan int
	on bool
}


func(tv *SonyTv) turnOn(){
	if !tv.on{
		tv.on = true
	}
}

func(tv *SonyTv) volumeUp(){
	tv.currentVol += 1
}

func(tv *SonyTv) volumeDown(){
	tv.currentVol -= 1
}

func(tv *SonyTv) setChannel(channel int){
	tv.currentChan = channel
}