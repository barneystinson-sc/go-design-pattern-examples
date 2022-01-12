package Iterator

import "fmt"

func IteratorCallBackExample(){
	library_obj.IterateBooks(BookInfo)

	//passing callback function

	library_obj.IterateBooks(func(book Book) error{
		fmt.Println("Author is",book.Author)
		return nil
	})

	// Using iterator

	lib_iterator := createIterator()

	for true{
		if lib_iterator.hasNext(){
			fmt.Println(lib_iterator.next())
		}else{
			break
		}
	}
}