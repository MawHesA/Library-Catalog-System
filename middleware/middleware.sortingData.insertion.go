package middleware

import "librarycatalog/global"

func SortingDataInsertion(BookData *global.Data, n int) {

	var i, j int
	var temp global.Book

	for i = 1; i < n; i++ {

		temp = BookData[i]
		j = i - 1

		for j >= 0 && BookData[j].Publishyear > temp.Publishyear {
			BookData[j+1] = BookData[j]
			j = j - 1
		}

		BookData[j+1] = temp
	}
}
