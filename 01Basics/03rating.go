package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main(){

	var name string 
	var userRating string

	//Frontend
	//take name as input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(("Please enter ur full name:"))
	name, _ = reader.ReadString('\n')

	// take rating from user and convert it to int to float
	reader = bufio.NewReader((os.Stdin))
	fmt.Println(("Please rate our dosa center between 1 and 5: "))
	userRating,_ = reader.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	//Backend
	// fmt.Printf("%v, %v", name, mynumRating)
	fmt.Printf("Hello %v, \n Thanks for rating our dosa center with %v star rating. \n\n Your rating was recorder in our system at %v",name, mynumRating, time.Now().Format(time.Stamp) )

	if mynumRating == 5 {
		fmt.Println("Bonus for team for 5 star service")
	} else if mynumRating == 4 || mynumRating == 3 {
		fmt.Println("We aew always improving")
	} else if mynumRating < 3 {
		fmt.Println("Need serious work on our side")
	}

	
}