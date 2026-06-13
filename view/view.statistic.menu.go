package view

import (
	"fmt"
	"librarycatalog/global"
)

func PrintStatisticMenu(stat global.Statistic) {
	fmt.Println("========== LIBRARY STATISTICS ==========")

	fmt.Println("\nTotal Books :", stat.TotalBooks)

	fmt.Println("\nBooks Per Category")
	fmt.Println("Fiksi      :", stat.Fiksi)
	fmt.Println("Non-Fiksi  :", stat.NonFiksi)
	fmt.Println("Romance    :", stat.Romance)
	fmt.Println("Horror     :", stat.Horror)
	fmt.Println("Fantasy    :", stat.Fantasy)

	fmt.Println("\nBooks Per Status")
	fmt.Println("Available   :", stat.Available)
	fmt.Println("UnAvailable :", stat.UnAvailable)
	fmt.Println("Borrowed    :", stat.Borrowed)

	fmt.Println("========================================")
}
