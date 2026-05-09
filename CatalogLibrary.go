package main
import("fmt"
		"librarycatalog/services"
)


func main(){
	var menu int

	fmt.Println("Menu")
	fmt.Println("1. Add Data")
	fmt.Println("2. Data Stored")
	fmt.Println("3. Update Existed Data")
	fmt.Println("4. Delete Existed Data")


	fmt.Print("Entered the desired number :")
	fmt.Scan(&menu)

	if menu == 1 {
		services.Create()
	}else if menu == 2 {
		read()
	}else if menu == 3 {
		update()
	}else{
		delete()
	}


}


func read(){
	fmt.Print("test")
}

func update(){

}

func delete(){

}
