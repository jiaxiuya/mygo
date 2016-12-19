package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total  // send total to c
}

// 无缓冲channel是在多个goroutine之间同步很棒的工具。
func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a) / 2], c)
	go sum(a[len(a) / 2:], c)
	x, y := <-c, <-c  // receive from c

	fmt.Println(x, y, x + y)

	c1 := make(chan int, 2)//修改2为1就报错，修改2为3可以正常运行
	c1 <- 1
	c1 <- 2
	fmt.Println(<-c1)
	fmt.Println(<-c1)

	c2 := make(chan int, 10)
	go fibonacci(cap(c2), c2)
	for i := range c2 {
		fmt.Println(i)
	}
	// 测试Channel是否已经关闭
	_, ok := <-c2
	if (ok) {
		fmt.Println("channel is close")
	}
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	close(c)
}