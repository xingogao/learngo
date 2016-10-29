package main

import (
	"sync"
	"bytes"
	"fmt"
)

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func main() {
	p := new(SyncedBuffer)
	var v SyncedBuffer
	fmt.Printf("%v\n%v\n", p.lock, v.buffer)
}
