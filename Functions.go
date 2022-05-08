package main

import (
	"fmt"
)

type greeter struct {
	greeting string
	name     string
}

func main() {
	//func is keyword for function
	//main is name of that function
	//() required syntax after main
	//
	fmt.Println("Hello, playground!")
	for i := 0; i < 5; i++ {
		sayMessage("Hello Go!", i)
	}
	greeting := "Hello"
	name := "Stacey"
	saygreeting(greeting, name)
	fmt.Println(name)
	s := sum(1, 2, 3, 4, 5) //variadic parameters
	fmt.Println(*s)         //pointer s changed to dereferenced pointer which prints value
	s1 := sum1(1, 2, 3, 4, 5)
	fmt.Println(s1)
	//anonymous function
	func() { //anonymous function doesn't need a function name
		msg := "Hello, Harsha!" // this msg vriable belongs to this function itself only
		//we also cannot access the variables which is inside the main function
		fmt.Println(msg)
	}() //invoke function () meaning it runs and executes at the same time
	//if you want to use outer variables inside anonymous function, do this!
	for i := 0; i < 5; i++ {
		/*we can also write f:=*/ var f func(i int) = func(i int) { //we can assign functions to variables but (i) changes to variable(i) in lower/new line
			fmt.Println(i)
		}
		f(i) //passing outer i inside the anonymous function
	}
	g := greeter{greeting: "Hello", name: "Hii"}
	g.greet()
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

func sayMessage(msg string, idx int) {
	fmt.Println(msg, idx) //msg is a local variable
}

func saygreeting(greeting, name string) {
	fmt.Println(greeting, name)
	name = "Ted"      //changing here doesn't change in main function
	fmt.Println(name) //however if you change te sending parameters addresses, then changing name into "Ted" changes in main()
}

//returning requires returning datatype before {} in function
func sum(values ...int) *int { //variadic parameters which makes parameters as a slice
	fmt.Println(values)
	result := 0
	for _, v := range values { //_ is theindexes of slice
		result += v
	}
	return &result //sending address of result which makes s as a pointer
}

func sum1(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return //function implisitly returns result int declared above in the function as (result int)
}

func divide(a, b float64) (float64, error) {
	//we do this because 0.0 returns +Inf meaning infinitely positive value
	if b == 0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
		//panic("Cannot provide zero as a second value")
	}
	return a / b, nil //when no error was present
}

func (g greeter) greet() { //method is a function that's executing in unknown context
	//the difference between method and a norma function is method uses (g greeter) in ()
	//greet() gets a copy of greeter under the name of g
	//(g greeter) greeter is valuetype not a pointer.. thus called value reciever which is a copy of struct greeter
	fmt.Println(g.greeting, g.name)
	//changing value reciever's/struct variable info inside this function, it doesn't change in main()
	//however using (g *greeter) without changing anything else in this function and then changing it changes the info in main() too..
}
