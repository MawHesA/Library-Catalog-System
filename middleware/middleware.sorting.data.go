package middleware

import "librarycatalog/global"

func Sortingdata(BookData *global.Data, n int) {

	var i, j, min int
	var temp global.Book

	for i = 0; i < n-1; i++ {

		min = i

		for j = i + 1; j < n; j++ {

			if BookData[j].BookId < BookData[min].BookId {

				min = j
			}
		}

		temp = BookData[i]
		BookData[i] = BookData[min]
		BookData[min] = temp
	}
}
