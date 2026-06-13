package view

import (
	"fmt"
)

func PrintStatisticMenu() {
	fmt.Println("========== LIBRARY STATISTICS ==========")

	fmt.Println("\nTotal Books :", n)

	fmt.Println("\nBooks Per Category")
	fmt.Println("Fiksi      :", fiksi)
	fmt.Println("Non-Fiksi  :", nonFiksi)
	fmt.Println("Romance    :", romance)
	fmt.Println("Horror     :", horror)
	fmt.Println("Fantasy    :", fantasy)

	fmt.Println("\nBooks Per Status")
	fmt.Println("Available   :", available)
	fmt.Println("UnAvailable :", unavailable)
	fmt.Println("Borrowed    :", borrowed)

	fmt.Println("========================================")
}
