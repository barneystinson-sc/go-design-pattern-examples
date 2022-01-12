package Factory

import "fmt"

type magazine struct{
	publication
}


func(nw *magazine) String(){
	fmt.Println("%s is a magazine of %s publication",nw.name,nw.publisher)
}

func createMagazine(name string,publisher string, pages int) IPublication{
	return &magazine{
		publication:publication{
			name:name,
			publisher: publisher,
			pages: pages,
		},
	}
}