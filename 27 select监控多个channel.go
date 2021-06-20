// package main

// import "fmt"

// func fibonacii(c, quit chan int) {
// 	x, y := 1, 1

// 	for {
// 		// select可以完成监控多个channel的状态
// 		select {
// 		case c <- x:
// 			// 如果c可写, 则该case就会进来
// 			fmt.Println("write:", x)
// 			x = y
// 			y = x + y
// 		case <-quit:
// 			fmt.Println("quit")
// 			return
// 		}
// 	}
// }

// func main() {
// 	c := make(chan int)
// 	quit := make(chan int)

// 	go func() {
// 		for i := 0; i < 6; i++ {
// 			fmt.Println("read:", <-c)
// 		}

// 		quit <- 0
// 	}()

// 	// channel作为形参, 传递的是引用
// 	fibonacii(c, quit)
// }
