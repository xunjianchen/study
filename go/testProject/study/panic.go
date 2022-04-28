package main

import (
	"fmt"
)

func divide() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Runtime panic caught: %v\n", err)
		}
	}()

	var i = 1
	var j = 0
	k := i / j
	fmt.Printf("%d / %d = %d\n", i, j, k)
}

func main() {
	divide()
	fmt.Println("divide方法调用完毕，回到main函数")
}
