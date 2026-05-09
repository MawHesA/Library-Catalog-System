package main
import(
	"fmt"
)

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
		fmt.Print(create())
	}else if menu == 2 {
		fmt.Print(read())
	}else if menu == 3 (
		fmt.Print(update())
	)else{
		fmt.Print(delete())
	}


}

func create(){

}

func read(){

}

func update(){

}

func delete(){

}