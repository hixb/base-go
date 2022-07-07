package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
for, if 后面的条件没有括号
if 条件里可以定义变量
没有while
switch 不需要 break, 也可以直接switch多个条件
*/

// 进制
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

// 读取文件内容
func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 死循环
func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),         // 101
		convertToBin(13),        // 1011
		convertToBin(345345435), // 10100100101011000110110011011
		convertToBin(0),         // 0
	)
	readFile("abc.txt")
	//forever()
}
