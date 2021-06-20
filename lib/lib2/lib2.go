package lib2

import "fmt"

// 当前lib2包提供的接口
func Lib2Test() {
	fmt.Println("API from lib2")
}

func init() {
	fmt.Println("lib2 init")
}
