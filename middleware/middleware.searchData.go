package middleware

import (
	"fmt"
	"librarycatalog/global"
)

func Search(BookData *global.Data, n int, idx *int) {
	var i int
	var target string
	var target2 int
	var menu string
	var left,right int
	var loc,temp int

	*idx = -1

	fmt.Println("Type 'back' to return or press any number to continue")

	fmt.Scan(&menu)
	if menu == "back" {
		return
	}

	fmt.Println("============================================")
	fmt.Print("Please Enter The Title : ")
	fmt.Scan(&target)
	fmt.Print("Please Enter the BookId : ")
	fmt.Scan(&target2)

	fmt.Println("================================================================")
	fmt.Printf("%-10s %-20s %-15s %-15s %-10s %-12s\n",
				"BOOK ID",
				"TITLE",
				"CATEGORY",
				"WRITTER",
				"YEAR",
				"STATUS")


	for i = 0; i < n; i++ {

		if BookData[i].Title == target {
			temp = temp + 1
		}
	}

	if temp > 1 {
	
	left  = 0
	right = n -1
	loc = -1

	for left <= right && loc == -1 {
		mid := (left + right)/2

		if BookData[mid].BookId == target2 {
			loc = mid
		}else if BookData[mid].BookId < target2 {
			left = mid + 1
		}else{
			right = mid - 1
		}

		*idx = loc
	}
	}else{
		for i = 0; i < n; i++ {

		if BookData[i].Title == target {
			*idx = *idx + 1
			fmt.Printf("%-10d %-20s %-15s %-15s %-10d %-15s\n",
				BookData[*idx].BookId,
				BookData[*idx].Title,
				BookData[*idx].Category,
				BookData[*idx].Writter,
				BookData[*idx].Publishyear,
				BookData[*idx].Status)

			}
		}
	}


	if temp > 1 {
	fmt.Printf("%-10d %-20s %-15s %-15s %-10d %-15s\n",
				BookData[loc].BookId,
				BookData[loc].Title,
				BookData[loc].Category,
				BookData[loc].Writter,
				BookData[loc].Publishyear,
				BookData[loc].Status)
	}
	

	if *idx == -1 {
		fmt.Println("Data Not Found")
		return
	}
}
