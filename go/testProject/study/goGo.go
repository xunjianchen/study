package main

import (
	"fmt"
	"time"
)

func aaa() {
	for {
		time.Sleep(time.Millisecond*200)
		fmt.Println("-1-1-1-1-1-1-1-")
	}
}

func main() {
	// go 协程的嵌套
	go func() {
		fmt.Println("----------1")
		go aaa()
		fmt.Println("---------2")
		return
	}()

	/*for {
		runtime.GC()
	}*/
}