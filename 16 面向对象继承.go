package main

// import "fmt"

// type Human struct {
// 	name string
// 	sex  string
// }

// func (human *Human) Eat() {
// 	fmt.Println("human eat")
// }

// func (human *Human) Walk() {
// 	fmt.Println("human walk")
// }

// type SuperMan struct {
// 	Human // SuperMan类继承了Human类的方法
// 	level int
// }

// // 重定义父类的eat方法
// func (superman *SuperMan) Eat() {
// 	fmt.Println("superman eat")
// }

// // 子类的新方法
// func (superman *SuperMan) Fly() {
// 	fmt.Println("superman fly")
// }

// func (superman *SuperMan) Show() {
// 	fmt.Println(superman.name, superman.sex, superman.level)
// }

// func main() {
// 	h := Human{"zhangsan", "female"}
// 	h.Eat()
// 	h.Walk()

// 	// 1. 安装结构体的内容定义一个子类对象
// 	// s := SuperMan{Human{"lisi", "male"}, 100}
// 	// 2. 声明一个子类对象并为其赋予属性
// 	var s SuperMan
// 	s.name = "lisi"
// 	s.sex = "male"
// 	s.level = 100

// 	s.Eat()  // 子类的方法
// 	s.Walk() // 父类的方法
// 	s.Fly()  // 子类的方法
// 	s.Show()
// }
