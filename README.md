# Library Catalog System

A library catalog management system developed to manage book records, members, borrowing transactions, and book returns efficiently.

The application begins by displaying the main menu of the catalog system, allowing users to navigate through the available features. The menu consists of seven different options, each corresponding to a specific operation within the system.

The main menu provides the following functionalities:

1. Add New Book Data
2. Search Data by Published Year
3. Display Stored Books
4. Update Existing Book
5. Delete Existing Book
6. Book Statistics
7. Exit Program

After the menu is displayed, the system waits for the user to select one of the available options. The selected menu determines which function will be executed, and after the operation is completed, the program returns to the main menu until the user chooses to exit.

### 1. Add New Book Data (Create)

The first menu allows users to add new book records to the catalog. Users can input the required book information, and the system will store the data in the existing collection. The data will be automatically sorted base on book id after the data is being stored by the user.

func Create(BookData *global.Data, n *int) {

	var amount int
	var menu string

	validcategory := true
	validstatus := true
	validid:= true

	handler.PrintAddNewBookData()
	fmt.Println("Type 'back' to return or press any number to continue")

	fmt.Scan(&menu)
	if menu == "back" {
		return
	}

	fmt.Print("Enter Total Books To Be Added : ")
	fmt.Scan(&amount)

	for amount <= 0 {

		fmt.Println("Input must be more than 0")

		fmt.Print("Enter Total Books To Be Added : ")
		fmt.Scan(&amount)
	}

	fmt.Println()
	fmt.Println("======================================================")

	for i := 0; i < amount; i++ {
    validInput := true

		fmt.Printf("============ Data Input For Data Number %d ============\n",i+1)
    for validInput == true {

		fmt.Printf("Input Book %d Id : ", i+1)
        fmt.Scan(&BookData[*n].BookId)

		fmt.Printf("Input Book %d Title : ", i+1)
		fmt.Scan(&BookData[*n].Title)

		fmt.Printf("Input Book %d Category : ", i+1)
		fmt.Scan(&BookData[*n].Category)

		fmt.Printf("Input Book  %d Writter  : ", i+1)
		fmt.Scan(&BookData[*n].Writter)

		fmt.Printf("Input Book %d PublishedYear  : ", i+1)
		fmt.Scan(&BookData[*n].Publishyear)

		fmt.Printf("Input Book %d Status : ", i+1)
		fmt.Scan(&BookData[*n].Status)

		validcategory = 	BookData[*n].Category == "Fiksi" || 
							BookData[*n].Category == "Non-Fiksi" ||
							BookData[*n].Category == "Romance" ||
							BookData[*n].Category == "Horror" ||
							BookData[*n].Category == "Fantasy" 

		validstatus = 	BookData[*n].Status == "Available" ||
						BookData[*n].Status == "UnAvailable" ||
						BookData[*n].Status == "Borrowed"

		validid = BookData[*n].BookId > 0 

		if validid != true {
			 fmt.Println("============== The Input Must Be In Integer ==============")
		}
					
		if validcategory != true {
			fmt.Println()
			 fmt.Println("==============  Available category ==============")
			 fmt.Println("Fiksi | Non-Fiksi | Romance | Horror | Fantasy ")
			  fmt.Println("==============  Available category ==============")
		}

		if validstatus != true {
			fmt.Println()
			fmt.Println("================ Please Enter ================")
			fmt.Println("      Available | UnAvailable | Borrowed")
			fmt.Println("================ Wrong Status ================")
		}

		if validcategory && validstatus && validid == true {
			*n = *n + 1
			validInput = false
		}
		
    }
}

	middleware.Sortingdata(BookData, *n)
	handler.PrintAddedData()
}



### 2. Search Book Data

The second menu provides a search feature that enables users to find books based on specific information. In the current implementation, books can be searched using either the **Book ID** or the **Book Title**.

func Search(BookData *global.Data, n int, idx *int) {
	var i int
	var target int
	var menu string

	*idx = 1

	fmt.Println("Type 'back' to return or press any number to continue")

	fmt.Scan(&menu)
	if menu == "back" {
		return
	}
	fmt.Println("============================================")
	fmt.Print("Type the desired Year of publish book :")
	fmt.Scan(&target)

	for i = 0; i < n; i++ {

		if *idx == 1 {
			fmt.Println("========================== DATA FOUND ==========================")
			fmt.Printf("%-10s %-20s %-15s %-15s %-10s %-12s\n",
				"BOOK ID",
				"TITLE",
				"CATEGORY",
				"WRITTER",
				"YEAR",
				"STATUS")
			*idx = -1
		}
		if BookData[i].Publishyear == target {
			*idx = i
			fmt.Printf("%-10d %-20s %-15s %-15s %-10d %-15s\n",
				BookData[*idx].BookId,
				BookData[*idx].Title,
				BookData[*idx].Category,
				BookData[*idx].Writter,
				BookData[*idx].Publishyear,
				BookData[*idx].Status)
		}
	}

	if *idx == -1 {
		fmt.Println("Data Not Found")
		return
	}
}


### 3. Display Stored Books (Read)

The third menu displays all book records currently stored in the system. The application is initialized with several dummy records, allowing the display feature to function even if the user has not added any new data.

func Read(BookData *global.Data, n int) {

	for i := 0; i < n; i++ {
		fmt.Printf("%-10d %-20s %-15s %-15s %-10d %-15s\n",
			BookData[i].BookId,
			BookData[i].Title,
			BookData[i].Category,
			BookData[i].Writter,
			BookData[i].Publishyear,
			BookData[i].Status)
	}
	fmt.Print()
}

### 4. Update Existing Book

The fourth menu enables users to modify the information of an existing book. By selecting a stored record, users can update the book details with new information.

func Update(BookData *global.Data, n *int) {

	var target int
	var found int
	var menu string

	found = -1

	if *n == 0 {
		handler.PrintNoDataToUpdate()
		return
	}

	handler.PrintChooseDataToUpdate()

	Read(BookData, *n)
	fmt.Println("Type 'back' to return or press any number to continue")
	fmt.Scan(&menu)

	if menu == "back" {
		return
	}

	fmt.Print("Enter Book ID To Update : ")
	fmt.Scan(&target)
	fmt.Println()

	for i := 0; i < *n; i++ {
		if BookData[i].BookId == target {
			found = i
		}
	}

	if found == -1 {
		handler.PrintBookNotFound()
		return
	}

	for i := 0; i < *n; i++ {
		if BookData[i].BookId == target {
			fmt.Print("Enter New Title : ")
			fmt.Scan(&BookData[i].Title)
			fmt.Print("Enter New Category : ")
			fmt.Scan(&BookData[i].Category)
			fmt.Print("Enter New Writter : ")
			fmt.Scan(&BookData[i].Writter)
			fmt.Print("Enter New Publish Year : ")
			fmt.Scan(&BookData[i].Publishyear)
			fmt.Print("Enter New Status : ")
			fmt.Scan(&BookData[i].Status)
			break
		}
	}
}

### 5. Delete Existing Book

The fifth menu allows users to remove a book record from the catalog. The deletion process is performed by searching for the corresponding **Book ID** and removing the selected entry from the system.

func Delete(BookData *global.Data, n *int) {

	var target int
	var found int
	var confirmation, menu string

	found = -1

	if *n == 0 {

		handler.PrintNoDataToDelete()
		return
	}

	handler.PrintChooseDataToUpdate()
	Read(BookData, *n)

	fmt.Println("Type 'back' to return or press any number to continue")
	fmt.Scan(&menu)

	if menu == "back" {
		return
	}

	fmt.Print("Enter Book ID To Delete : ")
	fmt.Scan(&target)
	fmt.Println()

	for i := 0; i < *n; i++ {
		if BookData[i].BookId == target {
			found = i
		}
	}

	if found == -1 {

		handler.PrintBookNotFound()
		return
	}

	fmt.Println("Are you sure want to delete the data : y/n")
	fmt.Scan(&confirmation)

	for confirmation != "y" && confirmation != "Y" {

		if confirmation == "N" || confirmation == "n" {
			Delete(BookData, n)
			return
		} else {
			fmt.Println("Please Input either y or n!")
			fmt.Println("Are you sure want to delete the data : y/n")
			fmt.Scan(&confirmation)
		}
	}

	for i := found; i < *n-1; i++ {
		BookData[i] = BookData[i+1]
	}

	*n = *n - 1

	handler.PrintDeletedData()
}


### 6. Book Statistics

The sixth menu displays statistical information related to the library collection. The statistics include the availability status of books and the total number of borrowed books categorized by genre.

func Statistic(BookData *global.Data, n int) {

	var stat global.Statistic
	stat.TotalBooks = n

	for i := 0; i < n; i++ {

		switch BookData[i].Category {
		case "Fiksi":
			stat.Fiksi = stat.Fiksi + 1
		case "NonFiksi":
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


The program continues to run and return to the main menu after each operation until the user selects the **Exit Program** option, which terminates the applicationn.
