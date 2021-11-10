package main

import "fmt"

func main() {
	// make new
	// new - only allocates - no init of memory
	// make - allocate and init - non zeroed storage

	score := make(map[string]int)
	score["ananda"] = 24
	fmt.Println(score)

	score["aniket"] = 42
	score["subhranshu"] = 32
	score["ranjan"] = 21
	score["mishra"] = 40
	fmt.Println(score)

	getAnandaScore := score["ananda"]
	fmt.Println(getAnandaScore)

	delete(score, "aniket")
	fmt.Println(score)

	for k, v := range score {
		fmt.Printf("Score of %v is %v\n", k,v)

	}

}