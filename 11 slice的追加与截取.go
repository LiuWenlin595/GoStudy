package main

// import "fmt"

// func main() {
// 	num1 := make([]int, 3, 5)
// 	fmt.Println(num1, len(num1), cap(num1))

// 	// 向切片追加一个元素
// 	num1 = append(num1, 1)
// 	fmt.Println(num1, len(num1), cap(num1))

// 	// 向切片追加一个元素
// 	num1 = append(num1, 2)
// 	fmt.Println(num1, len(num1), cap(num1))

// 	// 超过cap容量, golang会在底层一次性扩张cap的容量
// 	num1 = append(num1, 3)
// 	fmt.Println(num1, len(num1), cap(num1))

// 	// 切片的变量名实际上是指向底层数组的指针
// 	s := []int{1, 2, 3}
// 	s1 := s[0:2]
// 	fmt.Println(s, s1)
// 	// s和s1的指针指向同一个位置, 所以s和s1的动态数组都会改变
// 	s1[0] = 100
// 	fmt.Println(s, s1)

// 	// 如果只想要一个数组改变需要进行copy深拷贝
// 	s2 := make([]int, len(s), cap(s))
// 	copy(s2, s)
// 	fmt.Println(s, s2)
// 	s2[0] = 50
// 	fmt.Println(s, s2)

// 	aa := [4]int{1, 2, 3, 4}
// 	// aa.append(aa, 5) // 定长数组不可以进行追加, 但是可以截取
// 	fmt.Println(aa[0:2])
// }
