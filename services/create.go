package services

import (
	"fmt"
	"librarycatalog/gdict"
)

func Create(BookData *gdict.Data, n *int) {

	var amount int
	var menu string

	fmt.Println("======================================================")
	fmt.Println("                 ADD NEW BOOK DATA")
	fmt.Println("======================================================")
	fmt.Println("Type 'back' to return or press any number to continue")
	
	fmt.Scan(&menu)
	if menu == "back"{
		return
	}

	fmt.Println("Input Format :")
	fmt.Println("BookId Title Category Writter Publishyear Status")
	fmt.Println("======================================================")

	fmt.Print("Enter Total Books To Be Added : ")
	fmt.Scan(&amount)

	for amount <= 0 {

	fmt.Println("Input must be more than 0")

	fmt.Print("Enter Total Books To Be Added : ")
	fmt.Scan(&amount)
}

	fmt.Println()
	fmt.Println("======================================================")
	
	for i := 0; i < amount; i++ {

		fmt.Printf("Input Book Data %-3d : ", i+1)

		fmt.Scan(
			&BookData[*n].BookId,
			&BookData[*n].Title,
			&BookData[*n].Category,
			&BookData[*n].Writter,
			&BookData[*n].Publishyear,
			&BookData[*n].Status,
		)

		*n = *n + 1
	}

	Sortingdata(BookData, *n)

	fmt.Println("======================================================")
	fmt.Println("     Book Data Successfully Added And Sorted")
	fmt.Println("======================================================")
}

func Sortingdata(BookData *gdict.Data, n int) {

	var i, j, min int
	var temp gdict.Book

	for i = 0; i < n-1; i++ {

		min = i

		for j = i + 1; j < n; j++ {

			if BookData[j].BookId < BookData[min].BookId {

				min = j
			}
		}

		temp = BookData[i]
		BookData[i] = BookData[min]
		BookData[min] = temp
	}
}