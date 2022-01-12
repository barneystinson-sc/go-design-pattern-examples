package Factory

import "fmt"

func FactoryExample(){

	pub1,err := newPublication("NEWSPAPER","hindu","bennet coleman",12)
	if err==nil{
		pubDetails(pub1)
	}
	pub2,err := newPublication("MAGAZINE","hindu","bennet coleman",12)
	if err==nil{
		pubDetails(pub2)
	}

}


func pubDetails(publication IPublication){
	fmt.Println(publication)
	fmt.Printf("type %T\n",publication)
}
