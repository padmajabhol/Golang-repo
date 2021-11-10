package main

import "fmt"

type User struct {
	Name  string
	Email string
	age   int
}

func main() {
	rob := User{"rob", "rob@lco.dev", 34}
	fmt.Printf("%+v\n", rob)
	fmt.Printf("%+v\n", rob.Email)

	var sam = new(User)
	sam.Name = "sam"
	sam.Email = "sam@lco.dev"
	sam.age = 22

	fmt.Printf("%+v\n",sam.Email)

	var tobby = &User{"tobby", "tobby@lco.dev", 22}
	fmt.Printf("%+v\n",tobby)

}