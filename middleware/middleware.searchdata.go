package middleware

import (
	"fmt"
	"librarycatalog/gdict"
)

func Search(BookData *gdict.Data, n int, idx *int) {
	var i int
	var target int
	var menu string

	*idx = 1

	fmt.Println("Type 'back' to return or press any number to continue")

	fmt.Scan(&menu)
	if menu == "back" {
		return
	}
	fmt.Println("============================================")
	fmt.Print("Type the desired Year of publish book :")
	fmt.Scan(&target)

	for i = 0; i < n; i++ {

		if *idx == 1 {
			fmt.Println("========================== DATA FOUND ==========================")
			fmt.Printf("%-10s %-20s %-15s %-15s %-10s %-12s\n",
				"BOOK ID",
				"TITLE",
				"CATEGORY",
				"WRITTER",
				"YEAR",
				"STATUS")
			*idx = -1
		}
		if BookData[i].Publishyear == target {
			*idx = i
			fmt.Printf("%-10d %-20s %-15s %-15s %-10d %-15s\n",
				BookData[*idx].BookId,
				BookData[*idx].Title,
				BookData[*idx].Category,
				BookData[*idx].Writter,
				BookData[*idx].Publishyear,
				BookData[*idx].Status)
		}
	}

	if *idx == -1 {
		fmt.Println("Data Not Found")
		return
	}
}
