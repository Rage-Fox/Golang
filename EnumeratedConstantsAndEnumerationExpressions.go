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

const (
	a = iota //iota is a counter used for enumerated constants
	b = iota
	c = iota + 10 //we can perform arithmetic, bitwise and bitshift operators here
	//squaring needs math package but 1<<iota can increase every value to it's 2 powers
	//this block is contained of constants as a pattern.. so if your constants are untyped, compiler applies the same thing iota to all
)

const (
	a2 = iota
)

const (
	_ = iota //underscore represents compiler that whaever the value we use here, the user doesn't use it anywhere inside the program
)

const (
	isadmin           = 1 << iota //1
	isheadquarters                //2
	canseefinancials              //4
	canseeafrica                  //8
	canseeeurope                  //16
	caneenorthamerica             //32
)

func main() {
	fmt.Printf("%v,%T\n", a, a)
	fmt.Printf("%v,%T\n", b, b)
	fmt.Printf("%v,%T\n", c, c)
	fmt.Printf("%v,%T\n", a2, a2) //iota value resets here
	var roles byte = isadmin | canseefinancials | canseeeurope
	fmt.Printf("%b\n", roles)                              //storing the values as a byte/binary
	fmt.Printf("Is admin? %v\n", isadmin&roles == isadmin) //bitmasking and comparing
	fmt.Printf("Is HQ? %v\n", isheadquarters&roles == isheadquarters)
}
