package main

import "fmt"

type Book struct {
	title    string
	pageNum  int
	isReaded bool
}

func booky(book Book) (string, int, bool) {
	book.title = "Little Prince"
	book.pageNum = 100
	book.isReaded = true
	return book.title, book.pageNum, book.isReaded
}
func main() {
	fmt.Println(booky(Book{}))
}
