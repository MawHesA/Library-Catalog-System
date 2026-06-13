package controller

import (
	"librarycatalog/global"
	"librarycatalog/view"
)

func Statistic(BookData *global.Data, n int) {

	var stat global.Statistic
	stat.TotalBooks = n

	for i := 0; i < n; i++ {

		switch BookData[i].Category {
		case "Fiksi":
			stat.Fiksi = stat.Fiksi + 1
		case "Non-Fiksi":
			stat.NonFiksi = stat.NonFiksi + 1
		case "Romance":
			stat.Romance = stat.Romance + 1
		case "Horror":
			stat.Horror = stat.Horror + 1
		case "Fantasy":
			stat.Fantasy = stat.Fantasy + 1
		}

		switch BookData[i].Status {
		case "Available":
			stat.Available = stat.Available + 1
		case "UnAvailable":
			stat.UnAvailable = stat.UnAvailable + 1
		case "Borrowed":
			stat.Borrowed = stat.Borrowed + 1
		}
	}

	view.PrintStatisticMenu(stat)
}
