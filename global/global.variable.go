package global

const Nmax = 1000

type Book struct {
	BookId      int
	Title       string
	Category    string
	Writter     string
	Publishyear int
	Status      string
}

type Data [Nmax]Book

type Statistic struct {
	TotalBooks int

	Fiksi    int
	NonFiksi int
	Romance  int
	Horror   int
	Fantasy  int

	Available   int
	UnAvailable int
	Borrowed    int
}
