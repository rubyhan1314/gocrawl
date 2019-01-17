package main

import (
	"regexp"
	"fmt"
)

func main() {
	const text = "My email is hanru723@163.com@haha.com"

	//re,err := regexp.Compile("hanru723@163.com")
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)

	match := re.FindString(text) //查找一个

	fmt.Println(match)

}
