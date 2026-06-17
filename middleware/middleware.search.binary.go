package middleware

import (
	"fmt"
	"librarycatalog/global"
)

func BinarySearchID(BookData *global.Data, n int, targetID int, idx *int) {
	var left, mid, right int

	left = 0
	right = n - 1

	*idx = -1

	for left <= right {

		mid = (left + right) / 2

		if BookData[mid].BookId == targetID {

			*idx = mid

			fmt.Println("================================================================")
			fmt.Printf("%-10s %-20s %-15s %-15s %-10s %-12s\n",
				"BOOK ID",
				"TITLE",
				"CATEGORY",
				"WRITTER",
				"YEAR",
				"STATUS")

			fmt.Printf("%-10d %-20s %-15s %-15s %-10d %-15s\n",
				BookData[mid].BookId,
				BookData[mid].Title,
				BookData[mid].Category,
				BookData[mid].Writter,
				BookData[mid].Publishyear,
				BookData[mid].Status)

			return
		}

		if targetID < BookData[mid].BookId {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if *idx == -1 {
		fmt.Println("Data Not Found")
		return
	}
}
