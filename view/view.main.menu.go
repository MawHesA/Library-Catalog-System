package view

import (
	"fmt"
)

func PrintMainMenu() {
	fmt.Println("======================================================")
	fmt.Printf("%35s\n", "LIBRARY CATALOG MANAGEMENT")
	fmt.Println("======================================================")
	fmt.Printf("%-5s %-35s\n", "1.", "Add New Books")
	fmt.Printf("%-5s %-35s\n", "2.", "Search Books")
	fmt.Printf("%-5s %-35s\n", "3.", "Display Stored Books")
	fmt.Printf("%-5s %-35s\n", "4.", "Update Existing Books")
	fmt.Printf("%-5s %-35s\n", "5.", "Delete Existing Books")
	fmt.Printf("%-5s %-35s\n", "6.", "Books Statistic")
	fmt.Printf("%-5s %-35s\n", "7.", "Exit Program")
	fmt.Println("======================================================")
}
