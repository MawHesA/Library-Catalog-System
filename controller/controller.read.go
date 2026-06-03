package controller

import (
	"fmt"
	"librarycatalog/gdict"
)

func Read(BookData *gdict.Data, n int) {

	for i := 0; i < n; i++ {
		fmt.Printf("%-10d %-20s %-15s %-15s %-10d %-15s\n",
			BookData[i].BookId,
			BookData[i].Title,
			BookData[i].Category,
			BookData[i].Writter,
			BookData[i].Publishyear,
			BookData[i].Status)
	}
	fmt.Print()
}
