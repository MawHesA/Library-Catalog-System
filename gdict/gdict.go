package gdict

const Nmax = 1000

type Book struct {
	BookId    int
	Title      string
	Category string
	Writter string
	Publishyear int
	Status string
	
}

type Data [Nmax]Book