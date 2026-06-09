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

	validcategory := true
	validstatus := true
	validid:= true

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

		fmt.Println("Input Format :")
		fmt.Println("BookId Title Category Writter Publishyear Status")
		fmt.Println("======================================================")
        fmt.Printf("Input Book Data %-3d : ", i+1)

    for validInput == true {
        fmt.Scan(&BookData[*n].BookId)
		fmt.Scan(&BookData[*n].Title)
		fmt.Scan(&BookData[*n].Category)
		fmt.Scan(&BookData[*n].Writter)
		fmt.Scan(&BookData[*n].Publishyear)
		fmt.Scan(&BookData[*n].Status)

		validcategory = 	BookData[*n].Category == "Fiksi" || 
							BookData[*n].Category == "Non-Fiksi" ||
							BookData[*n].Category == "Romance" ||
							BookData[*n].Category == "Horror" ||
							BookData[*n].Category == "Fantasy" 

		validstatus = 	BookData[*n].Status == "Available" ||
						BookData[*n].Status == "UnAvailable" ||
						BookData[*n].Status == "Borrowed"

		validid = BookData[*n].BookId > 0 

		if validid != true {
			 fmt.Println("============== The Input Must Be In Integer ==============")
		}
					
		if validcategory != true {
			fmt.Println()
			 fmt.Println("==============  Available category ==============")
			 fmt.Println("Fiksi | Non-Fiksi | Romance | Horror | Fantasy ")
			  fmt.Println("==============  Available category ==============")
		}

		if validstatus != true {
			fmt.Println()
			fmt.Println("================ Please Enter ================")
			fmt.Println("      Available | UnAvailable | Borrowed")
			fmt.Println("================ Wrong Status ================")
		}

		if validInput || validcategory || validid == false {
			fmt.Println()
			fmt.Printf("Input Book Data %-3d : ", i+1)
		}

		if validcategory && validstatus && validid == true {
			*n = *n + 1
			validInput = false
		}
		
    }
}

	middleware.Sortingdata(BookData, *n)
	handler.PrintAddedData()
}

//