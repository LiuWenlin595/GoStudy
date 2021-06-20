package main

// import "fmt"

// // 接口类型的父类, 本质是一个指针
// type AnimalIF interface {
// 	Sleep()
// 	GetColor() string // 获取动物的颜色
// 	GetType() string  // 获取动物的种类
// }

// // 具体的类, 只要实现了AnimalIF定义的函数, 就相当于继承了AnimalIF的接口
// type Cat struct {
// 	color string
// }

// func (cat *Cat) Sleep() {
// 	fmt.Println("cat sleep")
// }

// func (cat *Cat) GetColor() string {
// 	return cat.color
// }

// func (cat *Cat) GetType() string {
// 	return "Cat"
// }

// // 具体的类, 只要实现了AnimalIF定义的函数, 就相当于继承了AnimalIF的接口
// type Dog struct {
// 	color string
// }

// func (dog *Dog) Sleep() {
// 	fmt.Println("dog sleep")
// }

// func (dog *Dog) GetColor() string {
// 	return dog.color
// }

// func (dog *Dog) GetType() string {
// 	return "Dog"
// }

// func showAnimal(animal AnimalIF) {
// 	animal.Sleep() // 多态
// 	fmt.Println(animal.GetColor(), animal.GetType())
// }

// func main() {
// 	var animal AnimalIF // 接口的数据类型, 父类指针
// 	animal = &Cat{"Green"}
// 	animal.Sleep()

// 	animal = &Dog{"Yellow"}
// 	animal.Sleep()

// 	cat := Cat{"Red"}
// 	dog := Dog{"Blue"}
// 	showAnimal(&cat)
// 	showAnimal(&dog)
// }
