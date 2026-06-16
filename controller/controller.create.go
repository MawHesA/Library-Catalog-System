package controller

import (
	"fmt"
	"librarycatalog/global"
	handler "librarycatalog/handler/create"
	"librarycatalog/middleware"
)

func Create(BookData *global.Data, n *int) {

	var amount int
	var menu string
	var confirm string

	validcategory := true
	validstatus := true
	validid := true
	validyear := true

	handler.PrintAddNewBookData()
	fmt.Println("Type 'back' to return or press any number to continue")

	fmt.Scan(&menu)
	if menu == "back" {
		return
	}

	fmt.Print("Enter Total Books To Be Added : ")
	fmt.Scan(&amount)

	for amount <= 0 {
		fmt.Println("Input must be more than 0")
		fmt.Print("Enter Total Books To Be Added : ")
		fmt.Scan(&amount)
	}

	fmt.Println()
	fmt.Println("======================================================")

	for i := 0; i < amount; i++ {
		validInput := true

		fmt.Printf("============ Data Input For Data Number %d ============\n", i+1)

		for validInput == true {

			fmt.Printf("Input Book %d Id : ", i+1)
			fmt.Scan(&BookData[*n].BookId)

			fmt.Printf("Input Book %d Title : ", i+1)
			fmt.Scan(&BookData[*n].Title)

			fmt.Printf("Input Book %d Category : ", i+1)
			fmt.Scan(&BookData[*n].Category)

			fmt.Printf("Input Book %d Writter : ", i+1)
			fmt.Scan(&BookData[*n].Writter)

			fmt.Printf("Input Book %d PublishedYear : ", i+1)
			fmt.Scan(&BookData[*n].Publishyear)

			fmt.Printf("Input Book %d Status : ", i+1)
			fmt.Scan(&BookData[*n].Status)

			duplicateTitle := false
			confirm = "yes"

			for j := 0; j < *n; j++ {
				if BookData[j].Title == BookData[*n].Title {
					duplicateTitle = true
				}
			}

			if duplicateTitle {
				fmt.Println()
				fmt.Println("==========================================")
				fmt.Println("Book title already exists!")
				fmt.Print("Do you still want to add it? (yes/no) : ")
				fmt.Scan(&confirm)

				for confirm != "yes" && confirm != "no" {
					fmt.Println("Please input only 'yes' or 'no'.")
					fmt.Print("Do you still want to add it? (yes/no) : ")
					fmt.Scan(&confirm)
				}
			}

			validcategory =
				BookData[*n].Category == "Fiksi" ||
					BookData[*n].Category == "Non-Fiksi" ||
					BookData[*n].Category == "Romance" ||
					BookData[*n].Category == "Horror" ||
					BookData[*n].Category == "Fantasy"

			validstatus =
				BookData[*n].Status == "Available" ||
					BookData[*n].Status == "UnAvailable" ||
					BookData[*n].Status == "Borrowed"

			validid = BookData[*n].BookId > 0

			validyear = BookData[*n].Publishyear <= 2026

			if !validid {
				fmt.Println("============== Book ID must be greater than 0 ==============")
			}

			if !validcategory {
				fmt.Println()
				fmt.Println("============== Available Category ==============")
				fmt.Println("Fiksi | Non-Fiksi | Romance | Horror | Fantasy")
				fmt.Println("================================================")
			}

			if !validstatus {
				fmt.Println()
				fmt.Println("================ Available Status ================")
				fmt.Println("Available | UnAvailable | Borrowed")
				fmt.Println("==================================================")
			}

			if !validyear {
				fmt.Println("================ The Book Must Be Under 2026! ================")
			}

			if validcategory && validstatus && validid && validyear {

				if duplicateTitle == true {
					if confirm == "yes" {
						*n = *n + 1
						validInput = false
					}
					if confirm == "no" {
						fmt.Println("Please input the book data again.")
					}
				} else {
					*n = *n + 1
					validInput = false
				}
			}
		}
	}

	middleware.SortingInsertion(BookData, *n)
	handler.PrintAddedData()
}
