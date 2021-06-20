package main

// import "fmt"

// // 生命一种行的数据类型myint, 是int的一个别名
// type myint int

// // 定义一个结构体
// type Book struct {
// 	title string
// 	auto  string
// }

// func changeBook1(book Book) {
// 	// 形参会开辟一个新的空间
// 	book.title = "C++"
// }

// func changeBook2(book *Book) {
// 	// 指针传递, 直接在原地址上修改
// 	book.title = "C++"
// }

// func main() {
// 	var a myint = 10
// 	fmt.Println(a)
// 	fmt.Printf("%T\n", a)

// 	var book Book
// 	book.title = "Golang"
// 	book.auto = "zhangsan"
// 	fmt.Printf("%v\n", book)

// 	changeBook1(book)
// 	fmt.Printf("%v\n", book)

// 	changeBook2(&book)
// 	fmt.Printf("%v\n", book)
// }
