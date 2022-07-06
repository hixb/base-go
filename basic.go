package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	bb = true
	cc = "kkk"
)

// 定义空变量
func variableZeroVal() {
	var a int
	var s string

	fmt.Printf("%d %q\n", a, s)
}

// 变量赋予初始值
func variableInitialVal() {
	var a, b int = 3, 4
	var s string = "hello"

	fmt.Println(a, s, b)
}

// 变量类型推导
func variableTypeDeduction() {
	var a, b, c, d = 3, 4, true, "heha"
	fmt.Println(a, b, c, d)
}

// 变量简短
func variableShorter() {
	a, b, c, d := 3, 4, true, "heha"
	b = 5
	fmt.Println(a, b, c, d)
}

// 欧拉公式
func euler() {
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

// 三角形, 强制类型转换
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 定义常量
func consts() {
	const (
		filename = "hello.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func main() {
	variableZeroVal()
	variableInitialVal()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc)
	euler()
	triangle()
	consts()
}
