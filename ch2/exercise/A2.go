package main

import "fmt"

func main() {
	//var sum int
	// Q21
	/*
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}
	*/

	// Q22
	/*
	i := 0 // 定义循环变量
	I: // 定义标签
	fmt.Printf("%d\n", i)
	i++
	if i < 10 {
		goto I // 跳转回标签
	}
	*/

	// Q23
	/*
	var arr [10] int // Create an array with 10 elements
	for i := 0; i < 10; i++ {
		arr[i] = i // Fill it one by one
	}
	fmt.Printf("%v", arr) // With %v Go prints the type
	*/

	// Q31
	/*
	for i := 1; i <= 100; i++ {
		fizz := (i % 3 == 0)
		buzz := ( i % 5 == 0)
		if fizz {
			fmt.Print("Fizz")
		}
		if buzz {
			fmt.Print("Buzz")
		}
		if !fizz && !buzz {
			fmt.Print(i)
		}
		fmt.Println()
	}
	*/

	// Q41
	/*
	var str string
	for i := 0; i < 100; i++ {
		str += "A"
		fmt.Println(str)
	}
	*/

	// Q42
	/*
	str := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("String %s\nLength: %d, Runes: %d\n", str,
		len([]byte(str)), utf8.RuneCount([]byte(str)))
	fmt.Println(len(str), utf8.RuneCountInString(str))
	*/

	// Q43
	/*
	str := "asSASA ddd dsjkdsjs dk"
	res := []byte(str)
	res[4], res[5], res[6] = 'a', 'b', 'c'
	fmt.Println(str)
	fmt.Println(string(res))
	*/

	// Q44
	/*
	s := "foobar"
	a := []byte(s) // Again a conversion
	// Reverse a
	for i, j := 0, len(a) - 1; i < j; i, j = i + 1, j - 1 {
		a[i], a[j] = a[j], a[i] // Parallel assignment
	}
	fmt.Printf("%s\n", string(a)) // Convert it back
	*/

	// Q51
	var xs  []float64
	var avg float64
	sum := 0.0
	switch len(xs) {
	case 0:
		avg = 0
	default:
		for _, v := range xs {
			sum += v
		}
		avg = sum / float64(len(xs))
	}
	fmt.Println(avg)

}