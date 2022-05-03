//boolean types
//numeric types(integers,floating point,complex numbers)
//text types
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Strings")
	s := "this is a string"
	fmt.Printf("%v,%T\n", s[2], s[2])
	fmt.Printf("%v,%T\n", string(s[2]), s[2])
	//string byte represent UTF8 but rune represents UTF32
	//string concatenation can be done by string1+string2
	//we cannot change s[2]="u" aka byte as a string
	s1 := []byte(s)
	fmt.Printf("%v,%T\n", s1, s1)
	//rune is always a int32 type
	r := 'a' //var r rune = 'a'
	fmt.Printf("%v,%T\n", r, r)
	//Rune is also known as Unicode(we can print emoji ascii values as well)
	//ASCII defines 128 characters, identified by the code points 0–127. Unicode, a superset of ASCII, defines the codespace of 1,114,112 code points.
}
