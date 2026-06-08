package main

import (
	"fmt"
	"librarycatalog/controller"
	"librarycatalog/global"
	handlerInvalid "librarycatalog/handler/invalid"
	handlerRead "librarycatalog/handler/read"
	"librarycatalog/middleware"
	"librarycatalog/model"
	handlerView "librarycatalog/view"
)

func main() {

	var menu, idx int
	var BookData global.Data
	var n int

	model.LoadDummyData(&BookData, &n)
	handlerView.PrintMainMenu()

	fmt.Print("Enter Your Desired Menu : ")
	fmt.Scan(&menu)

	for menu != 6 {

		fmt.Println()

		switch menu {

		case 1:
			controller.Create(&BookData, &n)
		case 2:
			middleware.Search(&BookData, n, &idx)

		case 3:

			fmt.Printf("%-10s %-20s %-15s %-15s %-10s %-12s\n",
				"BOOK ID",
				"TITLE",
				"CATEGORY",
				"WRITTER",
				"YEAR",
				"STATUS")

			fmt.Println("---------------------------------------------------------------------------------------------")
			controller.Read(&BookData, n)

		case 4:
			handlerRead.PrintStoredBookData()
			controller.Update(&BookData, &n)

		case 5:
			controller.Delete(&BookData, &n)

		default:
			handlerInvalid.PrintInvalidMenu()
		}

		fmt.Println()
		handlerView.PrintMainMenu()
		fmt.Scan(&menu)
	}

	handlerView.PrintExitMenu()
}

//nyimpen branch baru doang
