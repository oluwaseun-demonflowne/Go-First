package arrays

import "fmt"

func Arrays() {
	user := []string{}
	var first_name string
	var last_name string
	var email string
	// asking for user input
	fmt.Println("what is your first name")
	fmt.Scanln(&first_name)
	fmt.Println("what is your last name")
	fmt.Scanln(&last_name)
	fmt.Println("what is your email")
	fmt.Scanln(&email)
	user = append(user, first_name,last_name)
	fmt.Println("Tis is the user", user)
	// fmt.Println("let us test this array function first")
}