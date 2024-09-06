package loop

import (
	"fmt"
	"strings"
)

func Loops() {
	user := []string{}
	var first_name string
	// var lis int = 10;
	// for i := range  lis {
	// 	fmt.Println(i)
	// }
	var last_name string
	var email string
	//  using for alone here makes it run infinitly
	// for {
	// asking for user input
		fmt.Println("what is your first name")
		fmt.Scanln(&first_name)
		fmt.Println("what is your last name")
		fmt.Scanln(&last_name)
		fmt.Println("what is your email")
		fmt.Scanln(&email)
		user = append(user, first_name,last_name,email)
		fmt.Println("Tis is the user", user)
	// fmt.Println("let us test this array function first")
	// }
	// print only first name
	firstNames := []string{}
	for _, user :=  range user {
		var namess = strings.Fields(user)
		fmt.Println(namess)
		firstNames = append(firstNames, namess[0])
	}
	fmt.Println(firstNames)
}