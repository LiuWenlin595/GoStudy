package main

// import "fmt"

// // 值传递, 并不会swap a b
// func swap1(a, b int) {
// 	var tmp int = a
// 	a = b
// 	b = tmp
// }

// // 地址传递, 会swap a b
// func swap2(pa, pb *int) {
// 	var tmp int = *pa
// 	*pa = *pb
// 	*pb = tmp
// }

// func main() {
// 	var a int = 10
// 	var b int = 20

// 	fmt.Println(a, b)
// 	swap1(a, b)
// 	fmt.Println(a, b)
// 	swap2(&a, &b)
// 	fmt.Println(a, b)

// 	var p *int = &a
// 	fmt.Println(&a)
// 	fmt.Println(p)

// 	// 二级指针, 指针的指针
// 	var pp **int = &p
// 	fmt.Println(&p)
// 	fmt.Println(pp)
// }
