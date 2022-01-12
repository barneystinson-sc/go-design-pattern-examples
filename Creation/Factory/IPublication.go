package Factory


type IPublication interface {
	SetName(name string)
	setPages(pages int)
	setPublisher(name string)
	getName() string
	getPages() int
	getPublisher() string
	}

type publication struct{
	name string
	pages int
	publisher string
}

func(pb *publication) SetName(name string){
	pb.name = name
}
func(pb *publication)  setPages(pages int){
	pb.pages = pages
}
func(pb *publication)  setPublisher(publisher string){
	pb.publisher = publisher
}
func(pb *publication)getName() string{
	return pb.name
}
func(pb *publication)getPages() int{
	return pb.pages
}
func(pb *publication)getPublisher() string{
	return pb.publisher
}
