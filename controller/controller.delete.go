package controller

import (
	"fmt"
	"librarycatalog/gdict"
)

func Delete(BookData *gdict.Data, n *int) {

	var target int
	var found int
	var confirmation, menu string

	found = -1

	if *n == 0 {

		fmt.Println("======================================================")
		fmt.Println("            THERE IS NO DATA TO DELETE")
		fmt.Println("======================================================")

		return
	}

	fmt.Println("======================================================")
	fmt.Println("        CHOOSE THE BOOK ID THAT WILL BE DELETED")
	fmt.Println("======================================================")

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

		fmt.Println("======================================================")
		fmt.Println("              BOOK DATA NOT FOUND")
		fmt.Println("======================================================")

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

	fmt.Println("======================================================")
	fmt.Println("         BOOK DATA SUCCESSFULLY DELETED")
	fmt.Println("======================================================")

}
