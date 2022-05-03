//boolean types
//numeric types(integers,floating point,complex numbers)
//text types
package main

import (
	"fmt"
)

func main() {
	/*var n bool = true//false
	fmt.Printf("%v,%T", n, n)*/
	//if we don't give any initialisation, boolean type stores "0" aka false
	n := 1 == 1
	m := 1 == 2
	fmt.Printf("%v,%T\n", n, n)
	fmt.Printf("%v,%T", m, m)
}
