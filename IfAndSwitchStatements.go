package main

import (
	"fmt"
	"math"
)

func main() {
	statepopulations := map[string]int{ //map[type of the key]type of value
		"california":   39250017,
		"texas":        27862596,
		"florida":      20612439,
		"new york":     19745289,
		"pennsylvania": 12802503,
	}
	if pop, ok := statepopulations["florida"]; ok {
		fmt.Println(pop)
	}
	//outside printing pop shows error
	number := 50
	guess := 30
	if guess < 1 || guess > 100 { //if the first condition is true then go doen't need to check others.. this process is called short circuiting
		fmt.Println("Guess must be between 1 and 100")
	} else {
		if guess < number {
			fmt.Println("Too Low")
		} else if guess > number {
			fmt.Println("Too High")
		}
		if guess == number {
			fmt.Println("You got it!")
		}
	}
	fmt.Println(!true)
	num := 0.123 //0.1 outputs same
	//if num==math.Pow(math.Sqrt(num),2)
	if math.Abs(num/math.Pow(math.Sqrt(num), 2)-1) < 0.001 {
		fmt.Println("Same")
	} else {
		fmt.Println("Different")
	}
	//switches doesn't need break at end because go has their functions implicitly
	switch i := 2 + 3; i { //switch 2{}
	case 1, 5, 10:
		fmt.Println("one or five or ten")
	case 2: //you cannot ave overlapping cases
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}
	j := 10
	switch {
	case j <= 10:
		fmt.Println("Yess")
		fallthrough //when two or more cases work
	case j <= 20:
		fmt.Println("Yupp")
	default:
		fmt.Println("Nooooo")
	}
	var k interface{} = 1 //interface{} takes any type of data.. we can assign any type of value to that variable
	switch k.(type) {
	case int:
		fmt.Println("Integer")
		//break
		fmt.Println("This will print too..")
	case float64:
		fmt.Println("Float64")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Another type")
	}
}
