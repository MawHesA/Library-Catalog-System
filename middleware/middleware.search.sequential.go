package middleware

import (
	"fmt"
	"librarycatalog/global"
)

func SequentialSearch(BookData *global.Data, n int, targetTitle string, idx *int) {
	var i int

	fmt.Println("================================================================")
	fmt.Printf("%-10s %-20s %-15s %-15s %-10s %-12s\n",
		"BOOK ID",
		"TITLE",
		"CATEGORY",
		"WRITTER",
		"YEAR",
		"STATUS")

	for i = 0; i < n; i++ {

		if BookData[i].Title == targetTitle {
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
