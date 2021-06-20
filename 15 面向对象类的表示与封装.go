package main

// import "fmt"

// // 类的首字母大写说明其他包也能够访问
// type Hero struct {
// 	// 如果类的属性首字母大写, 那么该属性能够对外访问, 否则为该类私有
// 	Name  string
// 	Ad    int
// 	Level int
// }

// /*
// func (hero Hero) GetName() {
// 	fmt.Println(hero.Name)
// }

// func (hero Hero) SetName(newName string) {
// 	// 值传递
// 	// hero是调用该方法的对象的一个拷贝, 并不会改变原来对象的name值
// 	hero.Name = newName
// }

// func (hero Hero) Show() {
// 	fmt.Println(hero.Name, hero.Ad, hero.Level)
// }
// */

// // 需要传指针进行地址传递, 这样才可以保证在原地址上做修改
// func (hero *Hero) GetName() {
// 	fmt.Println(hero.Name)
// }

// func (hero *Hero) SetName(newName string) {
// 	hero.Name = newName
// }

// func (hero *Hero) Show() {
// 	fmt.Println(hero.Name, hero.Ad, hero.Level)
// }

// func main() {
// 	hero := Hero{
// 		Name:  "zhangsan",
// 		Ad:    100,
// 		Level: 1,
// 	}

// 	hero.Show()

// 	hero.SetName("lisi")

// 	hero.Show()
// }
