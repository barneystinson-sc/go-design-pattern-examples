package Observer

import "fmt"

type observer interface{
	onUpdate(data string)
}

type DataListener struct{
	Name string
}

func(dl *DataListener) onUpdate(data string){
	fmt.Println("Data updated to ",dl.Name,"new data is",data)
}
