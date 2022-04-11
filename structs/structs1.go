package main

type Person struct {
	name  string
	age   int
	tasks Tasks
}

type Tasks struct {
	tMorning   string
	tAfternoon string
	tNight     string
}
