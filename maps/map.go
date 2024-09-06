package mapss

import "fmt"

// "fmt"
// "strings"

var user = make([]User, 0)
type User struct {
	first_name       string
	last_name        string
	email           string
}

func Maps() {
	for{
	 first_name, last_name, email := getUserInput()
	 var userss = User{
		first_name: first_name,
		last_name:  last_name,
		email:      email,
	 }
	 user = append(user, userss)
	 fmt.Println(user,"||", userss)
	first := printFirstName()
	fmt.Println(first)
	}
}

func getUserInput()(string, string,string){
	var first_name string
	var last_name string
	var email string
	fmt.Println("what is your first name")
	fmt.Scanln(&first_name)
	fmt.Println("what is your last name")
	fmt.Scanln(&last_name)
	fmt.Println("what is your email")
	fmt.Scanln(&email)
	return first_name, last_name, email
}	

func printFirstName() []string {
	firstNames := []string{}
	for _, use := range user {
		firstNames = append(firstNames, use.first_name)
	}
	return firstNames

}