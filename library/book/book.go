package book

import "fmt"

type Printable interface {
	PrintInfo()
}

func Print(p Printable) {
	p.PrintInfo()
}

type Book struct {
	title  string
	author string
	pages  int
}

func NewBook(title, author string, pages int) *Book {
	return &Book{
		title:  title,
		author: author,
		pages:  pages,
	}
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) SetPages(pages int) {
	b.pages = pages
}

func (b *Book) GetPages() int {
	return b.pages
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s, Author: %s, Pages: %d\n", b.title, b.author, b.pages)
}

type TextBook struct {
	Book
	grade int
}

func NewTextBook(title string, author string, pages int, grade int) *TextBook {
	return &TextBook{
		Book:  Book{title, author, pages},
		grade: grade,
	}
}

func (b *TextBook) PrintInfo() {
	fmt.Printf("Title: %s, Author: %s, Pages: %d, Grade: %d\n", b.title, b.author, b.pages, b.grade)
}
