package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	c := make(chan int, 3) // 带有缓冲的channel

// 	fmt.Println("len(c) =", len(c), ", cap(c) =", cap(c))

// 	go func() {
// 		defer fmt.Println("sub goroutine 结束...")

// 		// 因此channel的容量是3, 所以发送第三个元素后就会阻塞
// 		// 如果channel的容量是5, sub goroutine就会直接发送完五个, 结束任务
// 		for i := 0; i < 5; i++ {
// 			c <- i
// 			fmt.Println("sub goroutine, 发送的元素为:", i, "len(c) =", len(c), ", cap(c) =", cap(c))
// 		}
// 	}()

// 	time.Sleep(2 * time.Second)

// 	// 当main goroutine开始接收元素后, sub goroutine停止阻塞开始继续发数据
// 	for i := 0; i < 5; i++ {
// 		num := <-c // 从c中接受数据, 并赋值给num
// 		fmt.Println("main goroutine接收元素:", num)
// 	}
// 	fmt.Println("main goroutine 结束...")
// }
