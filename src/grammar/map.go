package main

import "fmt"

func main() {
	// 初始化一个字典
	rating := map[string]float32{"C":5, "Go":4.5, "Python":4.5, "C++":2 }
	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	delete(rating, "C")  // 删除key为C的元素

	// map也是一种引用类型，如果两个map同时指向一个底层，那么一个改变，另一个也相应的改变：
	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	m1 := m
	m1["Hello"] = "Salut"  // 现在m["hello"]的值已经是Salut了
}