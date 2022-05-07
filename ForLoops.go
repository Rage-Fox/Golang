package main

import (
	"fmt"
)

func main() {
	i := 0             // i scoped to main()
	for ; i < 5; i++ { //for i:=0;i<5;i++	//semicolon is compulsory even we don't do i:=0
		fmt.Println(i)
		if i%2 == 0 {
			i /= 2
		} else {
			i = 2*i + 1
		}
	}
	fmt.Println()
	//for i:=0,j:=0;i<5;i++,j++ is not allowed  because there is no , in go
	//i and j are scoped to that for loop
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 { //i++ is not a expression in go.. it is a statement. so we can't do i++,j++
		fmt.Println(i, j)
	}
	fmt.Println()
	i = 0
	for i < 5 { //removing both semicolons can work too..
		fmt.Println(i)
		i++
	}
	fmt.Println()
	i = 0
	for { //we can also write like this so that go knows were to stop the loop
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		} //if this condition satisfies.. go doesn't print the value down below
		fmt.Println(i)
	}
	fmt.Println()
Loop: //label
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop //telling go to break out to where label Loop exists
			}
		}
	}
	fmt.Println()
	s := []int{1, 2, 3}
	for key, value := range s { //key==index
		fmt.Println(key, value)
	} //this works for sices,arrays,maps,strings
	//if you only want value then put _ for key
}
