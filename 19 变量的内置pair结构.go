package main

// import "fmt"

// type Reader interface {
// 	Read()
// }

// type Writer interface {
// 	Write()
// }

// type Book struct {
// }

// func (book *Book) Read() {
// 	fmt.Println("read a book")
// }

// func (book *Book) Write() {
// 	fmt.Println("Write a book")
// }

// func main() {
// 	// 测试一
// 	// pair<static type:string, value:"abc">
// 	var a string = "abc"

// 	// pair<concrete type:, value:>, type实际上是interface所指向的具体类型
// 	var myString interface{}
// 	// pair<type:string, value:"abc">
// 	myString = a

// 	value, _ := myString.(string)
// 	fmt.Println(value)

// 	// 测试二
// 	// b: pair<type:Book, value:b的地址>
// 	b := &Book{}
// 	// r: pair<type:, value:>
// 	var r Reader
// 	// r: pair<type:Book, value:b的地址>
// 	r = b

// 	r.Read()

// 	// w: pair<type:, value:>
// 	var w Writer
// 	// w: pair<type:Book, value:b的地址>
// 	// 因为r是Book类型的, 可以断言为Writer
// 	// 断言成功之后把r赋值给w(也就是w得到了r的type和value, w也变成了Book)
// 	w = r.(Writer)

// 	w.Write()
// }
