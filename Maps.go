package main

import (
	"fmt"
)

func main() {
	//statepopulations:=make(map[string]int) and in below statement remove :
	statepopulations := map[string]int{ //map[type of the key]type of value
		"california":   39250017,
		"texas":        27862596,
		"florida":      20612439,
		"new york":     19745289,
		"pennsylvania": 12802503,
	}
	fmt.Println(statepopulations)
	//we cannot insert a slice into map like m:=map[[]int]string{} but can insert an array like m:=[[3]int]string{}
	fmt.Println(statepopulations["florida"])
	statepopulations["georgia"] = 10310371   //adding information into map
	fmt.Println(statepopulations)            //adding gerogia we cannot guarantee the order of map upon returning
	delete(statepopulations, "georgia")      //removing an element from map
	fmt.Println(statepopulations["georgia"]) //this shows 0 as if it is present in the map or it doesn't exist
	pop, ok := statepopulations["georgia"]   //ok shows true or false about whether it is present in the map or not
	fmt.Println(pop, ok)
	fmt.Println(len(statepopulations))
	sp := statepopulations     //manipulating map in on eplace impacts maps on every other place
	delete(sp, "pennsylvania") //works as a pointer(just to understand)
	fmt.Println(sp)
	fmt.Println(statepopulations)
}
