package main

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )

// // 子goroutine
// func newTask() {
// 	i := 0
// 	for {
// 		i++
// 		fmt.Printf("sub goroutine, i = %v\n", i)
// 		time.Sleep(1 * time.Second)
// 	}
// }

// // 主goroutine
// // 主goroutine退出后, 子goroutine也会退出
// func main() {
// 	// 创建一个go线程, 去执行newTask()流程
// 	// go newTask()

// 	// 用goroutine调用一个匿名函数
// 	// flag属于main goroutine, func函数属于sub goroutine, 所以想要传值flag := func()的话需要用channel
// 	go func(a, b int) bool {
// 		defer fmt.Println("A.defer")
// 		fmt.Printf("a = %v, b = %v\n", a, b)

// 		func() {
// 			defer fmt.Println("B.defer")

// 			// 退出当前goroutine, 而不是像return一样只退出当前函数
// 			runtime.Goexit()

// 			fmt.Println("B")
// 		}() // 后面需要加一个括号, 表示调用这个匿名函数

// 		fmt.Println("A")
// 		return true
// 	}(10, 20)

// 	i := 0
// 	for {
// 		i++
// 		fmt.Printf("main goroutine, i = %v\n", i)
// 		time.Sleep(3 * time.Second)
// 	}
// }
