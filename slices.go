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

// slice 扩展
func sliceExpand() {
	// s1的值为[2 3 4 5], s2的值为[5 6]
	// slice 可以向后扩展, 不可以向前扩展
	// s[i]不可以超越[len](s), 向后扩展不可哟超越底层数据cap(s)
	arr := [...]int{0, 1, 3, 4, 5, 6, 7}
	fmt.Println("arr =", arr)
	s1 := arr[2:6]
	s2 := s1[3:5] // [s1[3], s1[4]]

	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	var arr2 [5]int
	//array1()

	s1 := arr[2:]
	s2 := arr[:]

	fmt.Println("updateSlice(s1) ==============================")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("updateSlice(s2) ==============================")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("pointerPassing2 ==============================")
	pointerPassing2(arr2[:])
	fmt.Println(arr2)

	fmt.Println("Reslice ==============================")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("sliceExpand ==============================")
	sliceExpand()
}
