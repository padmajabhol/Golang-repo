package main

import "fmt"

func main() {
	var batman string = "I am batman"
	fmt.Println((batman))

	var superman string 
	superman = "I am superman"
	fmt.Println((superman))

	thor := "I am thor"
	fmt.Println((thor))

	// thorRating := 3.0
	// fmt.Printf("%v, %T", thorRating, thorRating) //T for data type

	var Ironman, CatAmerica string = "I am Ironman" , "I am Capt America"

	fmt.Println(Ironman, CatAmerica)

	var defaultValue string
	fmt.Println(defaultValue)

	var(
		spiderman = "I am spiderman"
		age = 18
		powers = "Web slings, spidey sense"
		antman  = "I am antman"
		)

		fmt.Println(spiderman, age, powers, antman)
}