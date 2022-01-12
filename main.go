package main

import (
	"design-pattern/Behavioral/Iterator"
	"design-pattern/Behavioral/Observer"
	"design-pattern/Creation/Builder"
	"design-pattern/Creation/Factory"
	"design-pattern/Creation/Singleton"
	"design-pattern/Structural/Adapter"
)

func main(){
	Builder.BuilderExample()
	Factory.FactoryExample()
	Singleton.SingletonExample()

	//Structural Patterns

	Adapter.AdapterExample()

	//Behavioral Patterns

	Observer.ObserverExample()
	Iterator.IteratorCallBackExample()
}