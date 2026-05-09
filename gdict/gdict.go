package gdict

const nmax = 1000

type Book struct {
	BookId    int
	Title      string
	Category string
	Writter string
	Publishyear string
	Status string
	
}

var BookData[nmax] Book