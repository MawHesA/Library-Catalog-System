package controller

import (
	"librarycatalog/global"
	"librarycatalog/view"
)

func Statistic(BookData *global.Data, n int) {

	var fiksi, nonFiksi, romance, horror, fantasy int
	var available, unavailable, borrowed int

	for i := 0; i < n; i++ {

		switch BookData[i].Category {
		case "Fiksi":
			fiksi = fiksi + 1
		case "Non-Fiksi":
			nonFiksi = nonFiksi + 1
		case "Romance":
			romance = romance + 1
		case "Horror":
			horror = horror + 1
		case "Fantasy":
			fantasy = fantasy + 1
		}

		switch BookData[i].Status {
		case "Available":
			available = available + 1
		case "UnAvailable":
			unavailable = unavailable + 1
		case "Borrowed":
			borrowed = borrowed + 1
		}
	}

	view.PrintStatisticMenu()
}
