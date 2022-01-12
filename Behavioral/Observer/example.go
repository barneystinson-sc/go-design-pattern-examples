package Observer

func ObserverExample(){
	Listener1 := DataListener{
		Name:"Listener 1",
	}
	Listener2 := DataListener{
		Name:"Listener 2",
	}

	dataSub := newDataSubjectInstance()

	dataSub.registerObserver(Listener1)
	dataSub.registerObserver(Listener2)

	dataSub.changeItem("New field")

	dataSub.unregisterObserver(Listener1)

	dataSub.changeItem("Unregistered Listener 1")
}
