package main

import (
	"fmt"
	"librarycatalog/controller"
	"librarycatalog/global"
	"librarycatalog/middleware"
	"librarycatalog/model"
)

func main() {

	var menu, idx int
	var BookData global.Data
	var n int

	model.LoadDummyData(&BookData, &n)

	fmt.Println("======================================================")
	fmt.Printf("%35s\n", "LIBRARY CATALOG MANAGEMENT")
	fmt.Println("======================================================")
	fmt.Printf("%-5s %-35s\n", "1.", "Add New Book Data")
	fmt.Printf("%-5s %-35s\n", "2.", "Search Data By Published Year")
	fmt.Printf("%-5s %-35s\n", "3.", "Display Stored Books")
	fmt.Printf("%-5s %-35s\n", "4.", "Update Existing Book")
	fmt.Printf("%-5s %-35s\n", "5.", "Delete Existing Book")
	fmt.Printf("%-5s %-35s\n", "6.", "Exit Program")
	fmt.Println("======================================================")

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
			fmt.Println("======================================================")
			fmt.Println("                  STORED BOOK DATA")
			fmt.Println("======================================================")

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
			fmt.Println("======================================================")
			fmt.Println("                UPDATE EXISTING BOOK")
			fmt.Println("======================================================")

			controller.Update(&BookData, &n)

		case 5:
			controller.Delete(&BookData, &n)

		default:
			fmt.Println("Invalid Menu Selection.")
			fmt.Println("Please Select The Available Menu Options.")
		}

		fmt.Println()
		fmt.Println("======================================================")
		fmt.Printf("%35s\n", "LIBRARY CATALOG MANAGEMENT")
		fmt.Println("======================================================")
		fmt.Printf("%-5s %-35s\n", "1.", "Add New Book Data")
		fmt.Printf("%-5s %-35s\n", "2.", "Search Data By Published Year")
		fmt.Printf("%-5s %-35s\n", "3.", "Display Stored Books")
		fmt.Printf("%-5s %-35s\n", "4.", "Update Existing Book")
		fmt.Printf("%-5s %-35s\n", "5.", "Delete Existing Book")
		fmt.Printf("%-5s %-35s\n", "6.", "Exit Program")
		fmt.Println("======================================================")
		fmt.Print("Enter Menu : ")
		fmt.Scan(&menu)
	}

	fmt.Println()
	fmt.Println("======================================================")
	fmt.Println("     Thank You For Using Library Catalog System")
	fmt.Println("======================================================")
}

//nyimpen branch baru doang
