package main
import"fmt"

const nmax = 1000

type Book struct {
	bookId    int
	title      string
	category string
	writter string
	publishyear string
	status string
	
}

var BookData[nmax] Book

func main(){
	var menu int

	fmt.Println("Menu")
	fmt.Println("1. Add Data")
	fmt.Println("2. Data Stored")
	fmt.Println("3. Update Existed Data")
	fmt.Println("4. Delete Existed Data")


	fmt.Print("Entered the desired number :")
	fmt.Scan(&menu)

	if menu == 1 {
		create()
	}else if menu == 2 {
		fmt.Print(read())
	}else if menu == 3 (
		fmt.Print(update())
	)else{
		fmt.Print(delete())
	}


}

func create(){
	var n int
	
	fmt.Println("Format : BookId,title,category,writter,publishyear,status")
	fmt.Print("Enter Desired Amounts Of Book to be changed :")
	fmt.Scan(&n)
	for i := 0 ; i < n ;i++{
	fmt.Scan(	&BookData[i].bookId,
				&BookData[i].title,
				&BookData[i].category,
				&BookData[i].writter,
				&BookData[i].publishyear,
				&BookData[i].status)
	}
}

func read(){
	fmt.Print("test")
}

func update(){

}

func delete(){

}
