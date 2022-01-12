package Adapter

type mitvAdapter struct{
	mitv *MiTV
}

func(tv *mitvAdapter) turnOn(){
	tv.mitv.turnOnTV()
}

func(tv *mitvAdapter) volumeUp(){
	tv.mitv.volumeUpTV()
}

func(tv *mitvAdapter) volumeDown(){
	tv.mitv.volumeDownTV()
}

func(tv *mitvAdapter) setChannel(channel int){
	tv.mitv.gotoChannel(channel)
}
