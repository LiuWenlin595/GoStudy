package main

// import "fmt"

// // 固定长度的数组是值传递
// func printArray1(myArray [4]int) {
// 	fmt.Println("print arrayFunction1")
// 	for index, value := range myArray {
// 		fmt.Println(index, value)
// 	}
// 	myArray[0] = 0
// }

// // 动态数组的是引用传递
// func printArray2(myArray []int) {
// 	fmt.Println("print arrayFunction2")
// 	for index, value := range myArray {
// 		fmt.Println(index, value)
// 	}
// 	myArray[0] = 0
// }

// func main() {
// 	// 固定长度的数组
// 	var myArray1 [5]int

// 	myArray2 := [5]int{1, 2, 3}

// 	fmt.Println("print myArray1")
// 	for i := 0; i < len(myArray1); i++ {
// 		fmt.Println(i, myArray1[i])
// 	}

// 	fmt.Println("print myArray2")
// 	for index, value := range myArray2 {
// 		fmt.Println(index, value)
// 	}

// 	fmt.Printf("type myArray1 = %T\n", myArray1)
// 	fmt.Printf("type myArray2 = %T\n", myArray2)

// 	myArray3 := [4]int{11, 22, 33, 44}
// 	printArray1(myArray3)
// 	fmt.Println("print myArray3")
// 	for index, value := range myArray3 {
// 		fmt.Println(index, value)
// 	}

// 	// slice, 动态数组, 切片
// 	myArray4 := []int{111, 222, 333, 444}
// 	printArray2(myArray4)
// 	fmt.Println("print myArray4")
// 	for _, value := range myArray4 {
// 		fmt.Println(value)
// 	}

// 	fmt.Printf("%T\n", myArray1)
// 	fmt.Printf("%T\n", myArray1[:])
// 	fmt.Printf("%T\n", myArray4)
// }
