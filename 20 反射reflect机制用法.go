package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func reflectNum(arg interface{}) {
// 	fmt.Println(reflect.TypeOf(arg), reflect.ValueOf(arg))
// }

// func reflectStruct(arg interface{}) {
// 	// 获取input的type
// 	inputType := reflect.TypeOf(arg)
// 	fmt.Println(inputType)

// 	// 获取input的value
// 	inputValue := reflect.ValueOf(arg)
// 	fmt.Println(inputValue)

// 	// 通过inputType获取struct里面的字段
// 	for i := 0; i < inputType.NumField(); i++ {
// 		field := inputType.Field(i)
// 		value := inputValue.Field(i).Interface()
// 		fmt.Printf("%v %v: %v\n", field.Name, field.Type, value)
// 	}

// 	// 通过inputType获取struct里面的方法
// 	for i := 0; i < inputType.NumMethod(); i++ {
// 		method := inputType.Method(i)
// 		fmt.Printf("%v: %v\n", method.Name, method.Type)
// 	}
// }

// func reflectStarStruct(arg interface{}) {
// 	// 获取input的type
// 	inputType := reflect.TypeOf(arg)
// 	fmt.Println(inputType)

// 	// 获取input的value
// 	inputValue := reflect.ValueOf(arg)
// 	fmt.Println(inputValue)

// 	// 通过inputType获取*struct里面的方法
// 	for i := 0; i < inputType.NumMethod(); i++ {
// 		method := inputType.Method(i)
// 		fmt.Printf("%v: %v\n", method.Name, method.Type)
// 	}
// }

// type User struct {
// 	Id   int
// 	Name string
// 	Age  int
// }

// func (user User) Call1() {
// 	fmt.Println("user is called1")
// }

// func (user *User) Call2() {
// 	fmt.Println("user is called2")
// }

// func main() {
// 	var num float64 = 3.14
// 	reflectNum(num)

// 	user := User{
// 		Id:   1,
// 		Name: "haha",
// 		Age:  18,
// 	}
// 	// 这里注意引用含有的方法和指针含有的方法是不一样的
// 	// 指针的方法包含引用的方法(虽然我也不知道为什么)
// 	reflectStruct(user)
// 	reflectStarStruct(&user)
// }
