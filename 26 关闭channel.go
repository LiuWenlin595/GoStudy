package main

// import (
// 	"fmt"
// )

// func main() {
// 	// channel每次都要通过make创建, 如果通过var的话会生成一个nil channel, 无论收发都会阻塞
// 	c := make(chan int)

// 	go func() {
// 		for i := 0; i < 5; i++ {
// 			c <- i
// 		}

// 		// close可以关闭一个channel
// 		// 关闭channel后无法再发送数据
// 		close(c)
// 	}()

// 	for {
// 		// 如果不加close的话, channel数据读空后会报错, 因为系统监测不可能再有新的数据发送给channel
// 		// 关闭channel后, 可以继续从channel接收数据, 直到channel为空后ok变为false
// 		if data, ok := <-c; ok {
// 			fmt.Println(data)
// 		} else {
// 			break
// 		}
// 	}

// 	// // 使用range操作channel, 与上述代码等价, 同样的如果不关闭也会报错
// 	// for data := range c {
// 	// 	fmt.Println(data)
// 	// }

// 	fmt.Println("main goroutine 结束...")
// }
