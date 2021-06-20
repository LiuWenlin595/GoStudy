package main

// import "fmt"

// func main() {
// 	// 1. 声明slice是一个切片, 并且初始化, 默认值为1, 2, 3, 长度为3
// 	// 采用这种定义方式的话, 切片的容量=切片的长度
// 	slice1 := []int{1, 2, 3}
// 	fmt.Println(slice1, len(slice1), cap(slice1))

// 	// 2. 声明slice是一个切片, 但没有给slice分配空间
// 	// 如果不定义容量的话, 容量=长度
// 	var slice2 []int
// 	slice2 = make([]int, 3)
// 	fmt.Println(slice2, len(slice2), cap(slice2))

// 	// 3. 声明slice是一个切片, 同时给slice分配空间
// 	var slice3 []int = make([]int, 3, 5)
// 	fmt.Println(slice3, len(slice3), cap(slice3))

// 	// 4. 通过:=推导出slice是一个切片
// 	slice4 := make([]int, 3, 5)
// 	fmt.Println(slice4, len(slice4), cap(slice4))

// 	var slice5 []int
// 	fmt.Println(slice5, len(slice5), cap(slice5))
// 	if slice5 == nil {
// 		fmt.Println("slice5是nil切片")
// 	}

// 	// 空切片是指虽然动态数组为空, 但是指针已经创建好了
// 	slice6 := make([]int, 0, 0)
// 	fmt.Println(slice6, len(slice6), cap(slice6))
// 	if slice6 == nil {
// 		fmt.Println("slice6是nil切片")
// 	} else {
// 		fmt.Println("slice6是空切片")
// 	}

// }
