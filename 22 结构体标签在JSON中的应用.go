package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Movie struct {
// 	// 定义在json中的名字
// 	Title  string   `json:"title"`
// 	Year   int      `json:"year"`
// 	Price  int      `json:"rmb"`
// 	Actors []string `json:"actors"`
// }

// func main() {
// 	movie := Movie{"喜剧之王", 2000, 10, []string{"周星驰", "张柏芝"}}

// 	// 编码的过程  结构体 -> JSON
// 	jsonStr, err := json.Marshal(movie)
// 	if err != nil {
// 		fmt.Println("error occured 1")
// 		return
// 	}
// 	fmt.Printf("jsonStr = %s\n", jsonStr)

// 	// 解码的过程  jsonStr -> 结构体
// 	myMovie := Movie{}
// 	err = json.Unmarshal(jsonStr, &myMovie)
// 	if err != nil {
// 		fmt.Println("error occured 2")
// 	}
// 	fmt.Printf("%v\n", myMovie)
// }
