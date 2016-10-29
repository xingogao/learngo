package main

import "fmt"

func main() {
	var p *int
	fmt.Printf("%v\n", p)

	var i int
	p = &i
	fmt.Printf("%v\n", p)

	//p = &i
	*p = 8
	fmt.Printf("%v\n", *p)
	fmt.Printf("%v\n", i)
}