package main

import "fmt"

func main() {
	i := 10 // declare var
	p := &i // define pointer to i

	fmt.Println(*p)
}
