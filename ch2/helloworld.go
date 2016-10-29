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

	// array[n:m] 从 array 创建了一个 slice，具有元素 n 到 m-1
	a := [...] int{1, 2, 3, 4, 5}
	s1 := a[2:4]
	s2 := a[1:5]
	s3 := a[:]
	s4 := a[:4]
	s5 := s2[:]
	fmt.Println(a, s1, s2, s3, s4, s5)

	fmt.Println(throwsPanic(func() {
		panic("panic here")
	}))

	var arr [10] int
	arr[0] = 42
	arr[1] = 13
	fmt.Printf("The first element is %d\n", arr[0])

	a1 := [2][2] int{[2] int{1, 2}, [2] int{3, 4} }
	a2 := [2][2] int{[...] int{1, 2}, [...] int{3, 4} }
	// 错误的写法 a3 := [...][...]int{[...]int{1, 2}, [...]int{3, 4}}
	a3 := [2][2] int{{1, 2}, {3, 4} }
	fmt.Println(a1, a2, a3)

	const (
		m = 12
		n = 3
	)
	var array[m] int
	slice := array[0:n]
	fmt.Println(len(slice) == n, cap(slice) == m, len(array) == cap(array))

	// array[n:m] 从 array 创建了一个 slice，具有元素 n 到 m-1
	array1 := [...] int{1, 2, 3, 4, 5}
	slice1 := array1[2:4]
	slice2 := array1[1:5]
	slice3 := array1[:]
	slice4 := array1[:4]
	slice5 := s2[:]
	fmt.Println(slice1, slice2, slice3, slice4, slice5)

	sl0 := [] int{0, 0}
	sl1 := append(sl0, 2)
	sl2 := append(sl1, 3, 5, 7)
	sl3 := append(sl2, sl0...)
	fmt.Println(sl0, sl1, sl2, sl3)

	var a0 = [...] int{0, 1, 2, 3, 4, 5, 6, 7}
	var s = make([] int, 6)
	n1 := copy(s, a0[0:]) // n1 == 6, s == []int{0, 1, 2, 3, 4, 5}
	fmt.Println(n1, s)
	n2 := copy(s, s[2:]) // n2 == 4, s == []int{2, 3, 4, 5, 4, 5}
	fmt.Println(n2, s)

	monthdays := make(map[string]int)
	monthdays = map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31, // 逗号是必须的
	}
	fmt.Println(monthdays)
	fmt.Printf("%d\n", monthdays["Dec"])

	year := 0
	for _, days := range monthdays {
		year += days
	}
	fmt.Printf("Number of days in a year: %d\n", year)

	monthdays["Undecim"] = 30 // 添加一个月
	monthdays["Feb"] = 29 // 闰年时重写这个元素

	var value int
	var present bool
	value, present = monthdays["Jan"] // 如果存在， present 则有值 true
	// 或者更接近 Go 的方式
	v, ok := monthdays["Jan"] // “逗号 ok”形式
	fmt.Println(value, present, v, ok)

	delete(monthdays, "Mar") // 删除"Mar" 吧，总是下雨的月
	fmt.Println(monthdays)

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
