package main

import (
	"fmt"
)

type mystruct struct {
	foo int
}

func main() {
	a := 42
	b := a //just copied
	fmt.Println(a, b)
	a = 27
	fmt.Println(a, b)
	var c int = 42
	var d *int = &c
	fmt.Println(c, d)
	fmt.Println(&c, d) //address
	fmt.Println(c, *d) //asterisk prints value
	c = 80             //changing c changes d too..
	fmt.Println(c, *d)
	*d = 14 //changing pointer d changes c too..
	fmt.Println(c, *d)
	e := [3]int{1, 2, 3}
	f := &e[0]
	g := &e[1]                        //go doesn't allow math to use for pointers &a[1]-4
	fmt.Printf("%v %p %p\n", e, f, g) //%p prints value
	fmt.Println("Creating a Pointer")
	var m1 mystruct
	m1 = mystruct{foo: 42} //initialisation works like this but not in new function
	fmt.Println(m1)
	var m2 *mystruct
	m2 = &mystruct{foo: 42}
	fmt.Println(m2)
	var ms *mystruct
	//before initialising what is ms holding?
	fmt.Println(ms)    //pointer that you don't initialize is actually going to initialize for you.. thus holds the value nil
	ms = new(mystruct) //as initialisation is not allowed in new, it prints 0
	fmt.Println(ms)
	(*ms).foo = 42
	fmt.Println((*ms).foo) //dereferencing and setting that pointer to initialize
	ms.foo = 43
	fmt.Println(ms.foo) //compiler helps us without even dereferencing it
	h := []int{1, 2, 3} //slice works as a pointer.. but this doesn't work on arrays
	i := h              //slice itself contains a pointer pointing to 1t element of that slice
	h[1] = 42
	fmt.Println(h, i)
	//maps also has this property just like slices
}
