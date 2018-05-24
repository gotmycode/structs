package main

import "fmt"

//define properties of struc, create definition of type person
type person struct {
	firstName string
	lastName  string
}

func main() {
	//one of three ways to create a new struc, order of definition is relying on order of fields
	//alex := person{"Alex", "Anderson"}

	//2nd way -- can change order of fields type person struct
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
}
