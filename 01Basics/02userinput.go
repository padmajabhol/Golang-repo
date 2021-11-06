package main

import (
	// "fmt"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){

	// var name string = "padmaja"
	// var a, _ = fmt.Println(name)
	// fmt.Println(a)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter your full name: ")
	// myname, _ := reader.ReadString('\n')
	// fmt.Println((myname))

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your rating: ")
	myrating, _ := reader.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(myrating), 64)
	fmt.Println(mynumRating + 2)
}