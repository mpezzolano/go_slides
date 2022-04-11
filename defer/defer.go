package main

import "fmt"

func main() {
	defer fmt.Println("show this message at the end of function")
	fmt.Println("show some text and do something")
}
