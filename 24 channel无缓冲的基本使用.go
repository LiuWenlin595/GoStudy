package main

// import "fmt"

// func main() {
// 	// 定义一个无缓冲channel
// 	// channel可以保证两个goroutine的同步性
// 	c := make(chan int)

// 	go func() {
// 		defer fmt.Println("sub routine 结束...")

// 		fmt.Println("sub routine 正在运行...")

// 		c <- 666 // 将666发送给c
// 	}()

// 	// 如果main routine执行的快, 则数据没有发送方, main routine会阻塞直到sub routine开始发送(c <- 666)
// 	// 如果go routine执行的快, 则数据没有接收方, sub routine会阻塞直到main routine开始接收(num := <-c)
// 	num := <-c // 从c中接收数据, 并赋值给num

// 	fmt.Println("num = ", num)
// 	fmt.Println("main goroutine 结束...")
// }
