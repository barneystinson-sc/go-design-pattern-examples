package Adapter


func AdapterExample(){
	//Defining MiTV

	tv1:=MiTV{
		currentVol: 10,
		currentChan: 14,
		on: false,
	}
	tv2:=SonyTv{
		currentVol: 10,
		currentChan: 14,
		on: false,
	}

	// As sony tv implements tv interface so we can directly call its functions

	tv2.turnOn()
	tv2.volumeUp()
	tv2.setChannel(105)

	// MiTV can't be used directly so we use its adapter

	tv1Adapter := mitvAdapter{
		mitv: &tv1,
	}

	tv1Adapter.turnOn()
	tv1Adapter.volumeUp()
	tv1Adapter.setChannel(105)
}
