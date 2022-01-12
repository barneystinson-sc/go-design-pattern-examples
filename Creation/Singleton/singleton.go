package Singleton

import (
	"fmt"
	"sync"
)

type MyLogger struct{
	Loglevel string
}

func(lg *MyLogger) Log(msg string){
	fmt.Println("Log: ",msg)
}

var logger *MyLogger
var once sync.Once

func getInstance() *MyLogger{
	once.Do(func(){
		fmt.Println("Creating logger")
		logger =  &MyLogger{}
	})

	fmt.Println("Returning existing logger")
	return logger
}
