package Iterator

type iterator interface{
	hasNext() bool
	next() *Book
}

type BookIterator struct{
	current int
	Collection []Book
}

func(btr *BookIterator) hasNext() bool{
	if btr.current<len(btr.Collection){
		return true
	}else{
		return false
	}
}

func(btr *BookIterator) next() *Book{
	if btr.hasNext(){
		bk := &btr.Collection[btr.current]
		btr.current+=1
		return bk
	}else{
		return nil
	}
}

func createIterator() iterator{
	return &BookIterator{
		current: 0,
		Collection: library_obj.Collection,
	}
}