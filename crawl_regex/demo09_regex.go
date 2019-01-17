package main

import (
	"regexp"
	"fmt"
)

func main() {

	/*
	贪婪模式和非贪婪模式
	 */

	b1 := regexp.MustCompile(`\d{2,4}`).MatchString(`123c4abc`)
	fmt.Println(b1)

	b2 := regexp.MustCompile(`\d+`).MatchString(`13774kdf393`)
	fmt.Println(b2)

	b3 := regexp.MustCompile(`\d{2,4}?`).MatchString(`123c4abc`) //非贪婪模式，尽可能少匹配
	fmt.Println(b3)


	b4 := regexp.MustCompile(`\d+?`).MatchString(`13774kdf393`)
	fmt.Println(b4)



	s1 := "This is a number 234-245-236"
	//获取数字部分
	b5 := regexp.MustCompile(`(.+)(\d+-\d+-\d+)`).FindAllStringSubmatch(s1,-1)
	fmt.Println(b5)
	fmt.Println(b5[0][1])
	fmt.Println(b5[0][2])

	//非贪婪
	b6 := regexp.MustCompile(`(.+?)(\d+-\d+-\d+)`).FindAllStringSubmatch(s1,-1)
	fmt.Println(b6)
	fmt.Println(b6[0][1])
	fmt.Println(b6[0][2])



}
