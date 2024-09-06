package variable

import "fmt"

func Variables() {
	var myName = "oluwaseun"
	fmt.Println("my name is",	myName)

	var name string = "Emmanuel"
	fmt.Println("name =",name)

	userName := "admin"
	fmt.Println("name =",userName)

	var sum int
	fmt.Println("The sum is", sum)

	part1, other := 1,5
	fmt.Println("part 1 is", part1, "other is", other)
}