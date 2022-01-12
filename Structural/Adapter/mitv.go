package Adapter


type MiTV struct{
	currentVol int
	currentChan int
	on bool
}


func(tv *MiTV) turnOnTV(){
	if !tv.on{
	tv.on = true
	}
}

func(tv *MiTV) volumeUpTV(){
	tv.currentVol += 1
}

func(tv *MiTV) volumeDownTV(){
	tv.currentVol -= 1
}

func(tv *MiTV) gotoChannel(channel int){
	tv.currentChan = channel
}