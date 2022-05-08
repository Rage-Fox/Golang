package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	fmt.Println("middle")
	fmt.Println("end")
	fmt.Println()
	fmt.Println("start")
	defer fmt.Println("middle") //defer keyword executes any functions that are passed into it after the function finishes it's final statement, but before it actually returns
	fmt.Println("end")
	fmt.Println()
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")
	//defer works on printing last to first
	a := "hi"
	defer fmt.Println(a)
	a = "hello" //defer takes argument at the time defer is called bu not at where the function executes
}
