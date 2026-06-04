package model

import (
	"librarycatalog/global"
)

func LoadDummyData(BookData *global.Data, n *int) {

	dummy := []global.Book{
		{
			BookId:      101,
			Title:       "LaskarPelangi",
			Category:    "Novel",
			Writter:     "AndreaHirata",
			Publishyear: 2005,
			Status:      "Available",
		},
		{
			BookId:      102,
			Title:       "Bumi",
			Category:    "Fantasy",
			Writter:     "TereLiye",
			Publishyear: 2014,
			Status:      "Borrowed",
		},
	}

	*n = len(dummy)
	for i := 0; i < *n; i++ {
		(*BookData)[i] = dummy[i]
	}
}
