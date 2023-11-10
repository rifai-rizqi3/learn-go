package variable

import "fmt"

func Testvariable() {
	// var name string

	// name = "Rizqi"
	// fmt.Println(name)

	// name = "Rizqi Nur"
	// fmt.Println(name)

	// name := "Rizqi"
	// fmt.Println(name)

	// name = "Rizqi Nur"
	// fmt.Println(name)

	var (
		firstName  = "Dwi"
		middleName = "Nur"
		lastName   = "Farikhin"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
	fmt.Println(firstName + " " + middleName + " " + lastName)
}
