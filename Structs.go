package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	name string `required max: "100"` //tags makes us set a limit.. give back ticks
	//tags require a package names reflect
	origin string
}

type animal struct {
	name   string
	origin string
}

type bird struct {
	animal //embedding another string into this string
	speed  float32
	canfly bool
}

type doctor struct {
	number     int
	actorname  string
	episodes   []string
	companions []string
} //uppercase variables means the struct is export to the package.. so use lowercase start
//struct type gathers information together that are related to one group of information

func main() {
	adoctor := doctor{
		number:    3,
		actorname: "Lisa",
		//it's okay not to initialize episode(ignore that)
		companions: []string{
			"Harsha",
			"Lalisa",
		},
	} //positional syntax requires everything to be in correct order
	//that means the things that you write in adoctor has to be in the order according to struct type if you don't give them variable names but given info with just commas
	/*
		adoctor := doctor{
			3,
			"Lisa",
			[]string{
				"Harsha",
				"Lalisa",
			},

	*/
	fmt.Println(adoctor)
	fmt.Println(adoctor.actorname, adoctor.companions[1])
	//anonymous struct is when you declare inside main() function in a single line
	asinger := struct{ name string }{name: "Marshmello"} //very short lived data
	anothersinger := asinger
	anothersinger.name = "Lalisa Manoban"
	fmt.Println(asinger) //even though we copy and change the value, the values remain independent
	fmt.Println(anothersinger)
	//however if we use address operator, things change in both structs
	othersinger := &asinger //pointer to the struct
	othersinger.name = "Lisaaaa"
	fmt.Println(asinger)
	fmt.Println(othersinger)
	b := bird{}
	b.name = "Emu"
	b.origin = "Australia"
	b.speed = 48
	b.canfly = false
	/*
		b:=bird{
			animal: animal{name: "Emu", origin: "Australia"},//when we write this we have to tell explicitly about internal animal struct
			speed: 48,
			canfly: false,
		}
	*/
	fmt.Println(b)
	fmt.Println(b.name) //go is handling the request of name field automatically to the embedded animal type for us
	//animal and bird are of no relation and are independent exept the fact that one embeds another
	t := reflect.TypeOf(Animal{})     //getting type of structtype
	field, _ := t.FieldByName("name") //accessing name by fieldbyname
	fmt.Println(field.Tag)            //accessing tag
}
