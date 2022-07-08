package main

import "fmt"

/**
slice 本身没有数据, 是对底层array的一个view
*/

// 切片
func array1() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])
}

func updateSlice(s []int) {
	s[0] = 100
}

func pointerPassing2(arr []int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	var arr2 [5]int
	//array1()

	s1 := arr[2:]
	s2 := arr[:]

	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	pointerPassing2(arr2[:])
	fmt.Println(arr2)
}
