package main

import "fmt"

func array0() {
	var arr1 [5]int                 // 0 0 0 0 0
	arr2 := [3]int{1, 2, 3}         // 1 2 3
	arr3 := [...]int{2, 3, 4, 5, 7} // 2 3 4 5 7

	var grid [4][5]int // 四行五列 // [[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]]

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
}

// 常规循环数组
func arrayFor1() {
	arr3 := [...]int{2, 3, 4, 5, 7}
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
}

// 特殊写法
func arrayFor2() {
	arr3 := [...]int{2, 3, 4, 5, 7}
	for i := range arr3 {
		fmt.Println(arr3[i])
	}
}

// 带有下标
func arrayFor3() {
	arr3 := [...]int{2, 3, 4, 5, 7}
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
}

// 不带有下标
func arrayFor4() {
	arr3 := [...]int{2, 3, 4, 5, 7}
	// 可通过下划线省略变量, 不仅 range, 任何地方都可以通过_省略变量
	for _, v := range arr3 {
		fmt.Println(v)
	}
}

// 只要下标
func arrayFor5() {
	arr3 := [...]int{2, 3, 4, 5, 7}

	for i := range arr3 {
		fmt.Println(i)
	}
}

// 数组值类型
func printArray(arr [5]int) {
	// [10]int 和 [20]int 是不同类型
	// 调用 func f(arr [10]int) 会拷贝数组
	// 在go语言中一般不直接使用数组
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

// 指针传递可变数组
func pointerPassing(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	//arr2 := [3]int{1, 3, 4}
	arr3 := [...]int{15, 6, 7, 7, 9}
	pointerPassing(&arr1)
	pointerPassing(&arr3)
	fmt.Println(arr1, arr3)
}
