package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

/**
返回值类型写在最后面
可返回多个值
函数作为参数
没有默认值, 可选参数
*/

// 函数基础使用
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)
	}
}

// 带余除法
func div(a, b int) (q, r int) {
	return a / b, a % b
}

// 函数式编程写法
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

// 函数传入多值: 可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	if result, err := eval(13, 3, "X"); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(result)
	}
	q, r := div(13, 3)
	fmt.Println(q, r)
	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))
	fmt.Println(sum(1, 2, 3))
}
