package main

func main() {
	i := 0
	Here:   //这行的第一个词,以冒号结束作为标签,标签名是大小写敏感的。
	println(i)
	i++
	goto Here   //跳转到Here去
}
