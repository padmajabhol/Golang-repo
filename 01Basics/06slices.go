package main

import (
	"fmt"
	"sort"
)

func main(){
	var things = []string{"maleta", "ropa", "reloj"}
	fmt.Println(things)

	things = append(things, "bolso")
	fmt.Println(things)

	things = append(things[1:]) //will remove maleta
	fmt.Println(things)

	things = append(things[1:len(things)-1]) //will remove first and last index too
	fmt.Println(things)

	heros := make([]string,3 , 3)
	heros[0] = "subhranshu"
	heros[1] = "ranjan"
	heros[2] = "mishra"
	fmt.Println(heros)

	heros = append(heros, "ananda")
	fmt.Println(heros)
	fmt.Println(cap(heros))

	myints := []int{4, 7, 3, 2, 8}

	isSorted := sort.IntsAreSorted(myints)
	fmt.Println("Are ints sorted: ", isSorted)

	sort.Ints(myints)
	fmt.Println("sorted:",myints)



}