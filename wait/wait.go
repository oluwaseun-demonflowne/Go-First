package wait

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func Wait() {
	wg.Add(1)
	go hello()
	fmt.Println("hello world")
	wg.Wait()
}

func hello() {
	time.Sleep(5 * time.Second)
	fmt.Println("hello world haha")
	fmt.Println("#################")
	wg.Done()
}