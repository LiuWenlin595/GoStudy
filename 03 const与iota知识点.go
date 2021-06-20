package main

// import "fmt"

// // const 来定义枚举类型
// const (
// 	// 可以在const() 添加一个关键字iota, 每行的iota都会累加1, 第一行的iota默认值为0
// 	A = iota // iota = 0
// 	B        // iota = 1
// 	C        // iota = 2
// )

// const (
// 	a, b = iota + 1, iota + 2 // iota = 0
// 	c, d                      // iota = 1
// 	e, f                      // iota = 2

// 	g, h = iota * 2, iota * 3 // iota = 3
// 	i, j                      // iota = 4
// )

// func main() {
// 	// 常量(只读属性)
// 	const length int = 10

// 	fmt.Println(length)
// 	fmt.Println(A, B, C)
// 	fmt.Println(a, b, c, d, e, f)
// 	fmt.Println(g, h, i, j)

// 	// iota只能够配合const()一起使用
// 	// var a int = iota // 报错

// }
