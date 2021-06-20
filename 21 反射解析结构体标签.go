package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type resume struct {
// 	// 为属性绑定标签, 一个属性可以绑定多个标签, 中间用空格间隔
// 	Name string `info:"name" doc:"我的名字"`
// 	Sex  string `info:"sex"`
// }

// func findTag(str interface{}) {
// 	// 得到当前struct全部的元素
// 	// 如果传入的是*struct的话就用 t := reflect.TypeOf(str).Elem(), 详细地得去看文档
// 	t := reflect.TypeOf(str)
// 	for i := 0; i < t.NumField(); i++ {
// 		tagInfo := t.Field(i).Tag.Get("info")
// 		tagDoc := t.Field(i).Tag.Get("doc")
// 		fmt.Println(tagInfo, tagDoc)
// 	}
// }

// func main() {
// 	var re resume
// 	findTag(re)
// }
