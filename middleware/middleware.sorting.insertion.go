package middleware

import "librarycatalog/global"

func SortingInsertion(BookData *global.Data, n int) {

	var i, j int
	var temp global.Book

	for i = 0; i < n-1; i++ {

		temp = BookData[i]
		j = i - 1

		for j >= 0 && BookData[j].Publishyear > temp.Publishyear {
			BookData[j+1] = BookData[j]
			j = j - 1
		}

		BookData[j+1] = temp
	}
}
