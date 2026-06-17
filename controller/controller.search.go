package controller

import (
	"fmt"
	"librarycatalog/global"
	"librarycatalog/middleware"
)

func Search(BookData *global.Data, n int, idx *int) {
	var searchType int
	var menu string

	*idx = -1

	fmt.Println("Type 'back' to return or press any number to continue")

	fmt.Scan(&menu)
	if menu == "back" {
		return
	}

	fmt.Println("============================================")
	fmt.Println("1. Search By ID")
	fmt.Println("2. Search By Title")
	fmt.Print("Choose Search Type : ")
	fmt.Scan(&searchType)

	switch searchType {

	case 1:
		var targetID int

		fmt.Print("Enter Book ID : ")
		fmt.Scan(&targetID)
		middleware.SortingSearch(BookData, n)
		middleware.BinarySearchID(BookData, n, targetID, idx)

	case 2:
		var targetTitle string

		fmt.Print("Enter Book Title : ")
		fmt.Scan(&targetTitle)
		middleware.SortingSearch(BookData, n)
		middleware.SequentialSearch(BookData, n, targetTitle, idx)

	default:
		fmt.Println("Invalid Menu")
	}
}
