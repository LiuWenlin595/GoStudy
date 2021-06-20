package main

// import "fmt"

// func main() {
// 	// Map的定义是无序的, 因为底层用hash表存储
// 	// 1. 声明myMap是一种map类型, key是string, value是int
// 	var myMap1 map[string]int
// 	if myMap1 == nil {
// 		fmt.Println("myMap是一个nil map")
// 	}
// 	myMap1 = make(map[string]int, 10)
// 	myMap1["haha"] = 1
// 	fmt.Println(myMap1, len(myMap1))
// 	myMap1["xixi"] = 2
// 	fmt.Println(myMap1, len(myMap1))

// 	// 2. 直接使用:=, 也可以不指定map的长度
// 	myMap2 := make(map[int]string)
// 	myMap2[1] = "haha"
// 	fmt.Println(myMap2, len(myMap2))
// 	myMap2[2] = "xixi"
// 	fmt.Println(myMap2, len(myMap2))

// 	// 直接定义
// 	myMap3 := map[string]string{
// 		"one": "haha",
// 		"two": "xixi",
// 	}
// 	fmt.Println(myMap3, len(myMap3))
// }
