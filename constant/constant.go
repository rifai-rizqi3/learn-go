package constant

import (
	"fmt"
	"log"
)

func Testconstant() {
	const firstName string = "Rizqi"
	const lastName = "Rifai"

	fmt.Println(firstName)
	fmt.Println(lastName)
	log.Println(firstName)
	log.Println(lastName)

	// const (
	// 	firstName string = "Rizqi"
	// 	lastName         = "Rifai"
	// )

	// fmt.Println(firstName)
	// fmt.Println(lastName)

	//error
	// firstName = "Tidak Bisa Diubah"
	// lastName = "Tidak Bisa Diubah"
}
