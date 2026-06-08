package controller

import (
	"fmt"
	"librarycatalog/global"
	handler "librarycatalog/handler/delete"
)

func Delete(BookData *global.Data, n *int) {

	var target int
	var found int
	var confirmation, menu string

	found = -1

	if *n == 0 {

		handler.PrintNoDataToDelete()
		return
	}

	handler.PrintChooseDataToUpdate()
	Read(BookData, *n)

	fmt.Println("Type 'back' to return or press any number to continue")
	fmt.Scan(&menu)

	if menu == "back" {
		return
	}

	fmt.Print("Enter Book ID To Delete : ")
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

	fmt.Println("Are you sure want to delete the data : y/n")
	fmt.Scan(&confirmation)

	for confirmation != "y" && confirmation != "Y" {

		if confirmation == "N" || confirmation == "n" {
			Delete(BookData, n)
			return
		} else {
			fmt.Println("Please Input either y or n!")
			fmt.Println("Are you sure want to delete the data : y/n")
			fmt.Scan(&confirmation)
		}
	}

	for i := found; i < *n-1; i++ {
		BookData[i] = BookData[i+1]
	}

	*n = *n - 1

	handler.PrintDeletedData()
}
