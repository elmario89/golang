package main

import "fmt"

type Book struct {
	id     int
	name   string
	author string
}

type Library struct {
	books   []Book
	address string
}

func (p Library) printAddress() {
	fmt.Println(p.address)
}

func (b Book) printBookInfo() {
	fmt.Println("The book name is", b.name)
	fmt.Println("The author name is", b.author)
}

func (p Library) printAllBooks() {
	for i := range p.books {
		p.books[i].printBookInfo()
	}
}

func main() {
	var testBooks = []Book{{1, "War and peace", "Lev Tolstoy"}, {2, "Blokada", "Chakovskui"}}
	var testLibrary = Library{address: "Prospekt prosveshenia"}
	testLibrary.books = testBooks

	testLibrary.printAddress()
	testLibrary.printAllBooks()
}
