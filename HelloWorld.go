package main

// 导包方式1
//import "fmt"
//import "math"
// 导包方式2
import (
	"fmt"
	"math"
)
//变量列表，后面bool为变量类型
var a, b, c bool
//包含初始值的变量定义，每个变量对应一个初始值
var m, n int = 1, 2

//程序入口
func main() {
	var i int
	var x, y, z = true, false, "Hello"
	fmt.Println("Hello World")
	fmt.Println(math.Pi)
	fmt.Println(add(1, 2))
	fmt.Println(adds(2, 3))
	fmt.Println(swap("Hello", "World"))
	fmt.Println(split(6))
	fmt.Println(splits(6))
	fmt.Println(i, a, b, c)
	fmt.Println(m, n)
	fmt.Println(x, y, z)
}

//函数（1）
func add(x int, y int) int {
	return x+y
}
//函数（2）
func adds(x, y int) (int) {
	return x + y
}

//多个返回值
func swap(x, y string) (string, string) {
	return y, x
}

//命名返回值
func split(sum int) (x, y int) {
	x = sum - 1
	y = sum + 1
	return
}
func splits(sum int) (x, y int) {
	x = sum - 1
	y = sum + 1
	return y, x
}