//slices are similar to arrays with small changes
package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //slices have []
	b := a                                    //slices automatically take as a pointer
	b[1] = 5                                  //change in slice pointer b change in a too..
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("Length: %v\n", len(a))   //elements in slice
	fmt.Printf("Capacity: %v\n", cap(a)) //capacity of slice
	c := a[:]                            //slice od all elements
	d := a[3:]                           //slice from 4th element to the end
	e := a[:6]                           //slice first 6 elements
	f := a[3:6]                          //slice the 4t,5th,6th elements(3rd will not be there)
	//the above slicing can also work with array [...]
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	g := make([]int, 4) //generally has 2-3 arguments.. integer type and size of slice
	fmt.Println(g)
	fmt.Printf("Length: %v\n", len(g))
	fmt.Printf("Capacity: %v\n", cap(g))
	//g:=make([]int,length,capacity)//where capacity elements can be underlined elements(invisible/full size)
	h := []int{}
	fmt.Println(h)
	fmt.Printf("Length: %v\n", len(h))
	fmt.Printf("Capacity: %v\n", cap(h))
	h = append(h, 1) //adding an element
	fmt.Println(h)
	fmt.Printf("Length: %v\n", len(h))
	fmt.Printf("Capacity: %v\n", cap(h))
	h = append(h, 2, 3, 4, 5)
	fmt.Println(h)
	fmt.Printf("Length: %v\n", len(h))
	fmt.Printf("Capacity: %v\n", cap(h)) //capacity increases because go created a memory location for this slice and creates underlined elements
	h = append(h, []int{6, 7, 8}...)     //int shows error, but ... spreads the elements inside the slice of 6,7,8
	fmt.Println(h)
	fmt.Printf("Length: %v\n", len(h))
	fmt.Printf("Capacity: %v\n", cap(h)) //go creaes the capacity so we don't know how much it'll be
	i := append(a[:2], a[3:]...)         //0th upto 2nd position(exclusive) and 3rd til end
	fmt.Println(i)
	//but this changes the inside a slice because i works as a pointer
	fmt.Println(a)
}
