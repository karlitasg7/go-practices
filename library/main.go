package main

import "go-practices/library/book"

func main() {

	var myBook = book.NewBook("Book 1", "my author", 100)

	myBook.PrintInfo()

	myBook.SetTitle("Book 2")
	myBook.PrintInfo()

	var texbook = book.NewTextBook("Text Book 1", "my author", 100, 1)
	texbook.PrintInfo()

	book.Print(myBook)
	book.Print(texbook)

}
