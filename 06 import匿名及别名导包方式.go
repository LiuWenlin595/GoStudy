package main

// // Golang特殊语法, 如果只导包却不使用会报错
// // _ 匿名导入, name 别名导入, 可以只导包而不使用库
// // . 把目标包的全部方法直接导入当前包, 不建议使用. , 因为可能会存在lib1和lib2存在相同函数名的情况, 会引起歧义
// import (
// 	_ "GoStudy/lib/lib1"
// 	mylib2 "GoStudy/lib/lib2"
// 	// . "GoStudy/lib/lib2"
// )

// func main() {
// 	// lib1.Lib1Test()
// 	mylib2.Lib2Test()
// 	// Lib2Test()
// }
