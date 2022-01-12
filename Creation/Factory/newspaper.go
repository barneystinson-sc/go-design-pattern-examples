package Factory

import "fmt"

type newspaper struct{
	publication
}


func(nw *newspaper) String(){
	fmt.Println("%s is a newspaper of %s publication",nw.name,nw.publisher)
}

func createNewspaper(name string,publisher string, pages int) IPublication{
	return &newspaper{
			publication:publication{
				name:name,
				publisher: publisher,
				pages: pages,
			},
	}
}