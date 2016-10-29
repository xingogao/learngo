package main

import "fmt"

func main() {
	a := func(a ...int) string {
		fmt.Println("Hello")
		defer func() {
			if x := recover(); x != nil {
				fmt.Println(x)
			}
		}()
		return string(a[1])
	}
	a(1)
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("自定义panic")
	fmt.Printf("%T\n", a)

	/*
	var xs = map[int]func() int{
		1: func() int {
			return 10
		},
		2: func() int {
			return 20
		},
		3: func() int {
			return 30
		}, // 必须有逗号
			*/
	/* ... */
	/*

	}
	fmt.Println(xs[3]())
	*/
}