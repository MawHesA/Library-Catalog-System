package controller

import (
	"fmt"
	"librarycatalog/global"
	handler "librarycatalog/handler/update"
)

func Update(BookData *global.Data, n *int) {

	var target int
	var found int
	var menu string

	found = -1

	if *n == 0 {
		handler.PrintNoDataToUpdate()
		return
	}

	handler.PrintChooseDataToUpdate()

	Read(BookData, *n)
	fmt.Println("Type 'back' to return or press any number to continue")
	fmt.Scan(&menu)

	if menu == "back" {
		return
	}

	fmt.Print("Enter Book ID To Update : ")
	fmt.Scan(&target)
	fmt.Println()

	for i := 0; i < *n; i++ {
		if BookData[i].BookId == target {
			found = i
		}
	}

	if found == -1 {
		handler.PrintBookNotFound()
		return
	}

	for i := 0; i < *n; i++ {
		if BookData[i].BookId == target {
			fmt.Print("Enter New Title : ")
			fmt.Scan(&BookData[i].Title)
			fmt.Print("Enter New Category : ")
			fmt.Scan(&BookData[i].Category)
			fmt.Print("Enter New Writter : ")
			fmt.Scan(&BookData[i].Writter)
			fmt.Print("Enter New Publish Year : ")
			fmt.Scan(&BookData[i].Publishyear)
			fmt.Print("Enter New Status : ")
			fmt.Scan(&BookData[i].Status)
			break
		}
	}
}
