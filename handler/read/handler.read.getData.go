package handler

import (
	"fmt"
)

func PrintGetData() {
	fmt.Printf("%-10s %-20s %-15s %-15s %-10s %-12s\n",
		"BOOK ID",
		"TITLE",
		"CATEGORY",
		"WRITTER",
		"YEAR",
		"STATUS")

	fmt.Println("---------------------------------------------------------------------------------------------")
}
