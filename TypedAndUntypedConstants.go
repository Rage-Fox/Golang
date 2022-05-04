/*
naming convention
typed constants
untyped constants
enumerated constants
enumeration expressions
*/
package main

import (
	"fmt"
)

const a int16 = 27 //inner constant a shadows outer constant(type and value changes)

func main() {
	const myConst int = 42 //starting uppercase will declare tis constant externally.. so better to start off the starting letter lowercase
	// if we change myConst = 27 the compiler shows error because we cannot change the constant
	fmt.Printf("%v,%T\n", myConst, myConst)
	/*
		const myConst float 64 = math.Sin(1.57)
		compiler shows error because we can't set our constants equal to something that has to be determined at the runtime
	*/
	const (
		a int     = 14 //removing this prints the value and type of constant a outside main()
		b string  = "foo"
		c float32 = 3.14
		d bool    = true
	)
	fmt.Printf("%v,%T\n", a, a)
	fmt.Printf("%v,%T\n", b, b)
	fmt.Printf("%v,%T\n", c, c)
	fmt.Printf("%v,%T\n", d, d)
	var e int = 16
	fmt.Println(a + e) //we can add constant and variable if both are of same data types
	//if we give const a = 15 with no data type, we can perform operations with any type of variable with it!
	//the reason is compiler thinks const a as a literal and replaces a with 15 everywhere inside the program
}
