package main

import (
	"regexp"
	"fmt"
)

func main() {
	const text = `
My email is hanru723@163.com@haha.com
email is wangergou@sina.com
email is kongyixueyuan@cldy.org.cn
`

	//re,err := regexp.Compile("hanru723@163.com")
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)

	match := re.FindAllString(text,-1) //-1代表查找所有

	fmt.Println(match)

}
