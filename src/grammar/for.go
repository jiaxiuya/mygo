package main

import "fmt"

func main() {

	sum := 0;
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("sum is equal to ", sum)

	// 有些时候如果我们忽略expression1和expression3：
	sum = 1
	for ; sum < 1000; {
		sum += sum
	}

	// 其中;也可以省略，那么就变成如下的代码了，是不是似曾相识？对，这就是while的功能。
	for sum < 1000 {
		sum += sum
	}

	for index := 10; index > 0; index-- {
		if index == 5 {
			break // 或者continue
		}
		fmt.Println(index)
	}
	// break打印出来10、9、8、7、6
	// continue打印出来10、9、8、7、6、4、3、2、1

	maptest := map[string]string{"jijiji":"didididi"};
	// for配合range可以用于读取slice和map的数据：
	for k, v := range maptest {
		fmt.Println("map's key:", k)
		fmt.Println("map's val:", v)
	}

	for _, v := range maptest {
		fmt.Println("map's val:", v)
	}
}
