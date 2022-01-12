package Factory

import "fmt"

func newPublication(publicationType string,name string, publisher string, pages int) (IPublication,error){
	if publicationType=="NEWSPAPER"{
		return createNewspaper(name, publisher, pages),nil
	}else if publicationType=="MAGAZINE"{
		return createMagazine(name, publisher, pages),nil
	}else{
		return nil, fmt.Errorf("Incorrect publication type")
	}
}
