package main

import "fmt"

type contactInfo struct {
	email   string
	pinCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p *person) updateName(num string) {
	p.firstName = p.firstName + "_" + num
}

func main() {
	// sijo := person{
	// 	firstName: "SIJO",
	// 	lastName:  "SAM",
	// 	contactInfo: contactInfo{
	// 		email:   "sijosam1905@gmail.com",
	// 		pinCode: 12345,
	// 	},
	// }

	check()
	// sijo.updateName("192")

	// fmt.Printf("%+v", sijo)
}

func check() {
	name := []string{"hi", "hello", "sam"}
	update(name)
	fmt.Println(name)
}

func update(s []string) {

	s[0] = "Buhha"

}
