package lib1

import "fmt"

// 当前lib1包提供的接口
func Lib1Test() {
	fmt.Println("API from lib1")
}

func init() {
	fmt.Println("lib1 init")
}
