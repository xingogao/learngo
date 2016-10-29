package main

import "fmt"

func main() {
	/*
	rec(0)
	fmt.Println()
	*/

	/*
	fmt.Println(nextInt([]byte{'a', 'b', 'c'}, 1))
	*/

	/*
	a := []byte{'1', '2', '3', '4'}
	var x int
	for i := 0; i < len(a); {
		// 没有 i++
		x, i = nextInt(a, i)
		fmt.Println(x)
	}
	*/

	/*
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	*/

	fmt.Println(f())
}

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 0
}

/*
func nextInt(b [] byte, i int) (int, int) {
	x := 0
	// 假设所有的都是数字
	for ; i < len(b); i++ {
		x = x * 10 + int(b[i]) - '0'
	}
	return x, i
}
*/

/*
func rec(i int) {
	if i == 10 {
		return
	}
	rec(i + 1)
	fmt.Printf("%d ", i)
}
*/