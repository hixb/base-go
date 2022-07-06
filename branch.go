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

func main() {
	judge()
}
