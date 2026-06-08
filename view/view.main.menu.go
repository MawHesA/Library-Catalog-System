package view

import (
	"fmt"
)

func PrintMainMenu() {
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
}
