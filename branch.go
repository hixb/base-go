package main

import (
	"fmt"
	"io/ioutil"
)

// 判断条件
func judge() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

// switch
func eval(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong scope: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	judge()
	fmt.Println(
		eval(0),
		eval(59),
		eval(60),
		eval(82),
		eval(99),
		eval(100),
		eval(101),
	)
}
