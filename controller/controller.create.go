package controller

import (
	"fmt"
	"librarycatalog/global"
	handler "librarycatalog/handler/create"
	"librarycatalog/middleware"
)

func Create(BookData *global.Data, n *int) {

	var amount int
	var menu string

	handler.PrintAddNewBookData()
	fmt.Println("Type 'back' to return or press any number to continue")

	fmt.Scan(&menu)
	if menu == "back" {
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

	middleware.Sortingdata(BookData, *n)
	handler.PrintAddedData()
}
