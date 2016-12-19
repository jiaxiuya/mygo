package main

// 这个点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名，
// 也就是前面你调用的fmt.Println("hello world")可以省略的写成Println("hello world")
import (
	//. "fmt"
)

// 别名操作的话调用包函数时前缀变成了我们的前缀，即f.Println("hello world")
import (
	/// f "fmt"
)

// _操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数。
import (
	// "database/sql"
	// _ "github.com/ziutek/mymysql/godrv"
)

