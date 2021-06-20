package main

// import "fmt"

// func foo1(a string, b int) int {
// 	fmt.Println(a, b)
// 	c := 100
// 	return c
// }

// // 返回多个返回值, 匿名的
// func foo2(a string, b int) (int, int) {
// 	fmt.Println(a, b)
// 	return 100, 200
// }

// // 返回多个返回值, 有形参名称的
// func foo3(a string, b int) (r1 int, r2 int) {
// 	fmt.Println(a, b)
// 	// r1, r2 属于foo3的形参, 初始化默认的值是0
// 	// r1, r2 的作用域空间是 整个函数体foo3{}的空间
// 	fmt.Println(r1, r2)
// 	// 给有名称的返回值变量赋值
// 	r1, r2 = 1000, 2000
// 	return
// }

// func foo4(a string, b int) (r1, r2 int) {
// 	fmt.Println(a, b)
// 	fmt.Println(r1, r2)
// 	r1, r2 = 10000, 20000
// 	return
// }

// func main() {
// 	ret := foo1("abc", 111)
// 	fmt.Println(ret)

// 	ret1, ret2 := foo2("abc", 222)
// 	fmt.Println(ret1, ret2)

// 	ret1, ret2 = foo3("abc", 333)
// 	fmt.Println(ret1, ret2)

// 	ret1, ret2 = foo4("abc", 444)
// 	fmt.Println(ret1, ret2)
// }
