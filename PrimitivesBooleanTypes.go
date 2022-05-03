//boolean types
//numeric types(integers,floating point,complex numbers)
//text types
package main

import (
	"fmt"
)

func main() {
	/*
		n := 42           //int8,int16,int32,int64
		var m uint16 = 42 //uint8,uint16,uint32
		fmt.Printf("%v,%T\n", n, n)
		fmt.Printf("%v,%T", m, m)
	*/
	a := 10 //integer with integer
	b := 8
	fmt.Println("Arithmetic Operations")
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	var c int = 10
	var d int8 = 4
	fmt.Println(c + int(d)) //shows an error if we try to write c+d because without user's consent, go cannot perform operations on different data types
	// AND(&) OR(|) XOR(^) ANDNOT(&^)
	fmt.Println("Bitwise Operators")
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b) //neither one of the binary has to have bit for it to be true
	fmt.Println(b << 3)
	fmt.Println(b >> 3)
	//floating point in go follows IEEE-754 Standard--> float32,float64
	n := 3.14   //automatically saves to float64
	n = 13.7e72 //13.7 times to the 10 to the 72nd
	n = 2.1e14
	fmt.Printf("%v,%T\n", n, n)
	//var n float32 = 3.14 and while using this 13.7e72 shows error because it exceeds the limit
	//we can apply all arithmetic operetors except % on two float variables but bitwise and bitshift cannot be used on float datatypes in go
	var p complex64 = 1 + 2i
	fmt.Printf("%v,%T\n", p, p)
	fmt.Printf("%v,%T\n", real(p), real(p))
	fmt.Printf("%v,%T\n", imag(p), imag(p))
	//complex64 gives float32's
	var p1 complex128 = 2i
	fmt.Printf("%v,%T\n", p1, p1)
	fmt.Printf("%v,%T\n", real(p1), real(p1))
	fmt.Printf("%v,%T\n", imag(p1), imag(p1))
	//complex128 gives float64's
	//the same operations as for float are applied for complex datatype
	var n1 complex128 = complex(5, 12) //make other numbers as a complex using complex function
	fmt.Printf("%v,%T\n", n1, n1)
}
