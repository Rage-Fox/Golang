package main

import (
	"fmt"
)

func main() {
	a, b := 1, 0
	fmt.Println("start")
	defer fmt.Println("this was deferred") //panic happens after defer executes
	panic("something bad happened")        //we are not printing this but go prints this and panics and stops execution by showing error
	ans := a / b                           //it is invalid in go, so go panics and prints error
	fmt.Println("end")
	fmt.Println(ans)
} //recover function is also in control flows which recovers from panicking but it is a advanced topic
