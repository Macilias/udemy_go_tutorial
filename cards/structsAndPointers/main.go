package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	alex := person {firstName:"Alex", lastName:"Anderson"}
	//fmt.Print(alex)
	jim := person {
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 12345,
		},
	}

	//var alex person
	// print out all field names along with their values:
	fmt.Printf("%+v", alex)
	fmt.Println()

	jim.print()
	jim.updateName("Jon")
	jim.print()

	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()

	name := "Bill"

	fmt.Println(*&name)

	namePointer := &name

	fmt.Println(&namePointer)
	printPointer(namePointer)

}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}

func (p *person) updateName (newFirstName string) {
	p.firstName = newFirstName
}
