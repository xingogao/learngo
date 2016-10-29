package main

import (
	"fmt"
	"github.com/xingogao/learngo/ch4/even"
)

func main() {
	i := 5
	fmt.Printf("Is %d even? %v\n", i, even.Even(i))
}
