package main

import "fmt"

func main() {
	person1 := Person{
		name: "juan",
		age:  11,
		tasks: Tasks{
			tMorning:   "Wake up",
			tAfternoon: "Drink one Coffee",
			tNight:     "Sleep",
		},
	}

	fmt.Println(person1)

}
