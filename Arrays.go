package main

import (
	"fmt"
)

func main() {
	grades := [3]int{97, 85, 93} //[...] can also be written beside [3]
	fmt.Printf("Grades: %v\n", grades)
	var students [3]string //creates empty array
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa" //adding a element/string in array at some index
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #1: %v\n", students[0])
	fmt.Printf("Number of Students or size of the array: %v\n", len(students))
	var matrix [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	/*
		var matrix[3][3]int
		matrix[0]=[3]int{1,0,0}
		matrix[1]=[3]int{0,1,0}
		matrix[2]=[3]int{0,0,1}
	*/
	fmt.Println(matrix)
	a := [...]int{1, 2, 3}
	b := a //functions copies the entire array into another array in go
	b[1] = 5
	c := &b //pointer
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
