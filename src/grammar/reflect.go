package main

import (
	"reflect"
	"fmt"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	// 修改参数的值
	p := reflect.ValueOf(&x)
	s := p.Elem()
	s.SetFloat(7.1)
	println(x)
}
