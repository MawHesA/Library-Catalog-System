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

	handler.PrintAddNewBookData()
	fmt.Println("Type 'back' to return or press any number to continue")

	fmt.Scan(&menu)
	if menu == "back" {
		return
	}

	fmt.Println("Input Format :")
	fmt.Println("BookId Title Category Writter Publishyear Status")
	fmt.Println("======================================================")

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

    for validInput == true {
        fmt.Printf("Input Book Data %-3d : ", i+1)

        fmt.Scan(
            &BookData[*n].BookId,
            &BookData[*n].Title,
            &BookData[*n].Category,
            &BookData[*n].Writter,
            &BookData[*n].Publishyear,
            &BookData[*n].Status,
        )
		validcategory = 	BookData[*n].Category == "Fiksi" || 
							BookData[*n].Category == "Non-Fiksi" ||
							BookData[*n].Category == "Romance" ||
							BookData[*n].Category == "Horror" ||
							BookData[*n].Category == "Fantasy" 

		validstatus = 	BookData[*n].Status == "Available" ||
						BookData[*n].Status == "UnAvailable"
					

		if validcategory != true {
			fmt.Println()
			 fmt.Println("==============  Available category ==============")
			 fmt.Println("Fiksi | Non-Fiksi | Romance | Horror | Fantasy ")
			  fmt.Println("==============  Available category ==============")
		}

		if validstatus != true {
			fmt.Println()
			fmt.Println("===============================================")
			fmt.Println("Please enter either 'Available' or 'UnAvailable'")
			fmt.Println("===============================================")
		}

		if validcategory && validstatus  == true {
			*n = *n + 1
			validInput = false
		}
		
       
    }
}

	middleware.Sortingdata(BookData, *n)
	handler.PrintAddedData()
}

//