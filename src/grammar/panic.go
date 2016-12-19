package main

import (
	"os"
	"fmt"
)

func inits() {
	var user = os.Getenv("USER")
	if user == "" {
		panic("no value for $USER")
	} else {
		fmt.Println("has value for $USER")
	}
}

func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f() //执行函数f，如果f中出现了panic，那么就可以恢复回来'
	fmt.Println("recover")
	return
}

func main() {
	throwsPanic(inits)
}

