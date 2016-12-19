package main

import (
	"os"
	"fmt"
)

func main() {
	readFile()
}

func readFile() bool{

	file, err := os.Open("file")

	// 在defer后指定的函数会在函数退出前调用。
	defer file.Close()

	// 如果有很多调用defer，那么defer是采用后进先出模式，所以如下代码会输出4 3 2 1 0
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

	if file == nil {
		return false
	}
	if err != nil {
		return false
	}
	return true
}
