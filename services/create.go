package services

import ("fmt"
		"librarycatalog/gdict"

)

func Create(){

	var n int
	
	fmt.Println("Format : BookId,title,category,writter,publishyear,status")
	fmt.Print("Enter Desired Amounts Of Book to be changed :")
	fmt.Scan(&n)
	for i := 0 ; i < n ;i++{
	fmt.Scan(	&gdict.BookData[i].BookId,
				&gdict.BookData[i].Title,
				&gdict.BookData[i].Category,
				&gdict.BookData[i].Writter,
				&gdict.BookData[i].Publishyear,
				&gdict.BookData[i].Status)
	}
}