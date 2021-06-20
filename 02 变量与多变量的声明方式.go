package main

// import (
// 	"fmt"
// )

// /*
// 四种变量声明方式
// */

// // 声明全局变量: 1、2、3是可以的
// var gA int = 100
// var gB = 200

// // 方法四:= , 只能够用在函数体内来声明, 不能声明全局变量
// // 报错
// // gC := 300

// func main() {
// 	// 需要用printf来格式化输出, %T 可以打印变量的类型

// 	// 1. 声明一个变量, 默认值为0
// 	var a int
// 	fmt.Println(a)
// 	fmt.Printf("type a = %T\n", a)

// 	// 2. 声明一个变量, 初始化一个值
// 	var b int = 100
// 	fmt.Println(b)
// 	fmt.Printf("type b = %T\n", b)

// 	// 3. 初始化的时候, 可以省去数据结构, 通过值自动匹配当前的变量的数据类型
// 	var c = 100
// 	fmt.Println(c)
// 	fmt.Printf("type c = %T\n", c)

// 	var cc = "abcd"
// 	fmt.Println(cc)
// 	fmt.Printf("type cc = %T\n", cc)

// 	// 4. (常用的方法) 省去var方法, 直接自动匹配
// 	d := 100
// 	fmt.Println(d)
// 	fmt.Printf("type d = %T\n", d)

// 	// 全局变量
// 	fmt.Println(gA, gB)
// 	// fmt.Println(gC)   // 报错

// 	// 声明多变量
// 	var xx, yy int = 10, 20
// 	fmt.Println(xx, yy)
// 	var mm, nn = 30, "abcd"
// 	fmt.Println(mm, nn)

// 	// 多行的多变量声明
// 	var (
// 		ii int  = 100
// 		jj bool = true
// 	)
// 	fmt.Println(ii, jj)

// }
