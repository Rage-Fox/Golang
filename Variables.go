package main

import (
	"fmt"
	"strconv"
)

/*
we can give variable outside the main function
unused variables shows error in golang
more number of variables can be given like this
var (
	actorName string = "Elizabeth"
	doctorNumber int = 3
)
*/
var i int = 69

//lowercase variable means that variable belongs to that package itself
//uppercase variable are global
func main() {
	fmt.Println(i) //outside variable is printed first
	var i int
	i = 42
	//i:=27 cannot declare the same variable but can given as i=27 so that it can change
	var j float32 = 27.56
	k := 99.
	//printf means we can give and use %d formats
	fmt.Printf("%v,%T\n", i, i) //variable and datatype
	fmt.Printf("%v,%T\n", j, j)
	fmt.Printf("%v,%T\n", k, k)
	k = 990.90
	fmt.Println(k)
	i = int(j) //forcible format change(explicit conversion)
	fmt.Printf("%v,%T\n", i, i)
	i = 48
	fmt.Printf("%v,%T\n", i, i)
	//what happens if I convert integer to string, it doesnot prints the ascii value but instead prints the character in that ascii value
	//for that we need to use import strconv and use strconv.Itoa(i) to print the string ascii value as it is
	var l string
	l = string(rune(i))
	fmt.Printf("%v,%T\n", l, l)
	l = strconv.Itoa(i)
	fmt.Printf("%v,%T", l, l)
}
