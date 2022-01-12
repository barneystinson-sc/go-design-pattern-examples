package Iterator

import "fmt"

type BookType int

const(
	HardCover BookType = iota
	SoftCover
	PaperBack
	EBook
)

type Book struct{
	Name string
	Author string
	Publication string
	Pages int
}

type Library struct{
	Collection []Book
}

var library_obj Library = Library{
	Collection: []Book{{
		"Book1",
		"Author1",
		"Publication1",
		123,
		},
		{
			"Book2",
			"Author2",
			"Publication2",
			123,

		},
		{
			"Book3",
			"Author3",
			"Publication3",
			123,

		},
		{
			"Book4",
			"Author4",
			"Publication4",
			123,

		},
	},
}

func(bk *Book) String(){
	fmt.Println("Book ",bk.Name,"of publication",bk.Publication,"has",bk.Pages,"pages")}

func(lib *Library) IterateBooks(f func(Book) error){
	var err error
	for _,b := range lib.Collection{
		err = f(b)
		if err!=nil{
			fmt.Println("Error encountered :",err.Error())
		}
	}
}

func BookInfo(book Book) error{
	fmt.Println("Book ",book.Name,"of publication",book.Publication,"has",book.Pages,"pages")
	return nil
}
