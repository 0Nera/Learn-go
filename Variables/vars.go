package main

import "fmt"

func main() {
	var User1_Name string = "Aren"
	var User1_Age int = 16

	var (
		User2_Name string = "Tom"
		User2_Age  int    = 0
	)

	Login := "Tom2284055"
	Password := "328470354802"
	Year := 2022

	User2_Age = 27

	if Login == "Admin" && Password == "Root3921840282" {
		fmt.Println("Data:")
		fmt.Println(User1_Name, User1_Age, Year)
		fmt.Println(User2_Name, User2_Age, Year)
	} else {
		fmt.Println("Password incorrect")
	}
}
