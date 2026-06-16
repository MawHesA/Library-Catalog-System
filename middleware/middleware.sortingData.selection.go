package middleware

import "librarycatalog/global"

func SortingDataSelection(BookData *global.Data, n int) {

	var i, j, min int
	var temp global.Book

	for i = 0; i < n-1; i++ {

		min = i

		for j = i + 1; j < n; j++ {

			if BookData[j].Publishyear < BookData[min].Publishyear {
				min = j
			}
		}

		temp = BookData[i]
		BookData[i] = BookData[min]
		BookData[min] = temp
	}
}
