package arrays

import (
	"fmt"
	"sort"
)

func Arrays() {

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(ages)
	// sort directly mutates the array
	fmt.Println(ages)
	//  get the index of the number in the  int array
	index := sort.SearchInts(ages,30)
	fmt.Println(index)







	// user := []string{}
	// var first_name string
	// var last_name string
	// var email string
	// // asking for user input
	// fmt.Println("what is your first name")
	// fmt.Scanln(&first_name)
	// fmt.Println("what is your last name")
	// fmt.Scanln(&last_name)
	// fmt.Println("what is your email")
	// fmt.Scanln(&email)
	// user = append(user, first_name,last_name, email)
	// fmt.Println("Tis is the user", user)
	// // slice ranges
	// rangeOne := user[1:3]
	// fmt.Println(rangeOne)
}