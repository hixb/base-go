package main

import "fmt"

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

func main() {
	variableZeroVal()
	variableInitialVal()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc)
}
