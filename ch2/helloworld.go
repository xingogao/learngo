package main

import (
	"fmt"
) // 实现格式化的 I/O。

var (
	a = []rune{'a', 'b', 'c'}
)

const (
	b = iota
	c
)

/* Print something */
func main() {

	fmt.Printf("Hello, world; or κα λη µρα κ´ óσµ; orこんにちは世界\n")

	fmt.Println(a, b, c)

	var c complex64 = 5 + 5i; fmt.Printf("Value is: %v\n", c)

	fmt.Println(0 &^ 0, 0 &^ 1, 1 &^ 0, 1 &^ 1)

	if true && true {
		fmt.Println("true")
	}
	if ! false {
		fmt.Println("true")
	}

	test1()

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i // sum = sum + i 的简化写法
	}
	fmt.Println(sum)

	// Reverse a
	for i, j := 0, len(a) - 1; i < j; i, j = i + 1, j - 1 {
		// 平行赋值
		a[i], a[j] = a[j], a[i] // 这里也是
	}
	fmt.Println(a)

	for i := 0; i < 10; i++ {
		if i > 5 {
			break // 终止这个循环，只打印 0 到 5
		}
		fmt.Print(i)
	}
	fmt.Println()

	J: for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break J // 现在终止的是 j 循环，而不是 i 的那个
			}
			fmt.Print(i)
		}
	}
	fmt.Println()

	for i := 0; i < 10; i++ {
		if i > 5 {
			continue // 跳过循环中所有的代码println(i)
		}
		fmt.Print(i)
	}
	fmt.Println()

	list := [] string{"a", "b", "c", "d", "e", "f" }
	for k, v := range list {
		// 对 k 和 v 做想做的事情
		fmt.Print(k, v, ";")
	}
	fmt.Println()

	for pos, char := range "a Φx" {
		fmt.Printf("character '%c' starts at byte position %d\n", char, pos)
	}
	fmt.Println()

	fmt.Println(unhex('F'))

	switch 0 {
	case 0: // 空的 case 体
	case 1:
		fmt.Println("without fallthrough") // 当 i == 0 时， f 不会被调用！
	}
	switch 0 {
	case 0:
		fmt.Println("before fallthrough")
		fallthrough
	//fmt.Println("after fallthrough")
	case 1:
		fmt.Println("with fallthrough") // 当 i == 0 时， f 会被调用！
	}

	switch 2 {
	default:
		fmt.Println("default") // 当 i 不等于 0 或 1 时调用
		fallthrough
	case 0:
		fallthrough
	case 1:
		fmt.Println("case 1")
	}

	fmt.Println(shouldEscape(0))

	fmt.Println(Compare([]byte{0}, []byte{0}))

	fmt.Println(throwsPanic(func() {
		panic("panic here")
	}))

}

func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("recover()", x)
			b = true
		}
	}()
	f()
	return
}

// 比较返回两个字节数组字典数序先后的整数。
// 如果 a == b 返回 0，如果 a < b 返回 -1，而如果 a > b 返回 +1
func Compare(a, b [] byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	// 长度不同，则不相等
	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}
	return 0 // 字符串相等
}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+': // , as "or"
		return true
	}
	return false
}

func test1() error {
	var err error
	if err != nil {
		return err
	} else {
		return nil
	}
}
