package main

// import "fmt"

// // interface{}是空接口, 也是万能数据类型
// func myFunc(arg interface{}) {
// 	fmt.Println("myFunc is calling")
// 	// interface{} 如何区分此时引用的底层数据类型到底是什么?
// 	// 通过给interface{} 提供了"类型断言"机制
// 	// 以string为例
// 	value, ok := arg.(string)
// 	if !ok {
// 		fmt.Println(arg, "is not a string type")
// 	} else {
// 		fmt.Println(arg, "is a string type, value =", value)
// 		fmt.Printf("value type is %T\n", value)
// 	}
// }

// type Book struct {
// 	auto string
// }

// func main() {
// 	book := Book{"golang"}
// 	myFunc(book)
// 	myFunc("abc")
// 	myFunc(100)
// 	myFunc(3.14)
// }
