package uniq

import (
	"sync"
)

var (
	uniqLock sync.Mutex
	uniqId   int
)

func uniqFn() int {
	uniqLock.Lock()
	defer uniqLock.Unlock()
	uniqId++
	return uniqId
}

var uniq = make(chan int)

func init() {
	go func() {
		for i := 1; ; i++ {
			uniq <- i
		}
	}()
}
