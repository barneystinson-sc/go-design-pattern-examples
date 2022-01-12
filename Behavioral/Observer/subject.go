package Observer

type DataSubject struct{
	observers []DataListener
	field string
}

type observable interface{
	registerObserver(dl DataListener)
	unregisterObserver(dl DataListener)
	notifyAll()
}

func(ds *DataSubject) registerObserver(dl DataListener){
	ds.observers = append(ds.observers, dl)
}

func(ds *DataSubject) unregisterObserver(dl DataListener){
	newObservers := []DataListener{}
	for _,observer := range ds.observers{
		if observer.Name != dl.Name{
			newObservers = append(newObservers,observer)
		}
	}
	ds.observers = newObservers
}

func(ds *DataSubject) notifyAll(){
	for _,observer := range ds.observers{
		observer.onUpdate(ds.field)
	}
}

func(ds *DataSubject) changeItem(newfield string){
	ds.field = newfield
	ds.notifyAll()
}

func newDataSubjectInstance() *DataSubject{
	return &DataSubject{}
}


