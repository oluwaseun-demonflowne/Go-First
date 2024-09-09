package pointer

import "fmt"


//  means the value in that location is a string
func updateName(x *string) {
	//  get that pointer, see what it is pointing to then give it a new variable
	*x = "wedge"
}

func Pointer() {
	name := "tife"
	tpoint := &name
	updateName(&name)
	fmt.Println(name, *tpoint, &tpoint)
}