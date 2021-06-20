package main

// import "fmt"

// func deferFunc() int {
// 	fmt.Println("defer func")
// 	return 0
// }

// func returnFunc() int {
// 	fmt.Println("return func")
// 	return 0
// }

// func returnAndDefer() int {
// 	defer deferFunc()
// 	return returnFunc()
// }

// func add(a int) int {
// 	a += 1
// 	fmt.Println("in function: ", a)
// 	return a
// }

// func main() {
// 	// defer: 当前函数生命周期全部结束后, 才开始出栈
// 	// defer的执行顺序是栈, 先进后出
// 	// 但是defer中的计算会先执行
// 	defer fmt.Println("1")
// 	defer fmt.Println("2")
// 	fmt.Println("3")
// 	fmt.Println("4")

// 	returnAndDefer()

// 	var a int = 100
// 	fmt.Println("before defer: ", a)
// 	defer add(a)
// 	fmt.Println("after defer: ", a)
// }
