package main

import (
	"regexp"
	"fmt"
)

func main() {

	//符号：一般符号和特殊符号

	// 1..代表匹配任意字符
	b1 := regexp.MustCompile(`123.`).MatchString(`123*emeda`)
	fmt.Println(b1)

	b2 := regexp.MustCompile(`123....`).MatchString(`123memeda`)
	fmt.Println(b2)

	//2.[],匹配[]里的任意一个
	b3 := regexp.MustCompile(`[abc]`).MatchString(`saaa`) // 或a或b或c开头即可。
	fmt.Println(b3)
	b4 := regexp.MustCompile(`a[abc]`).MatchString(`as`) //第一个字母必须是a，第二个字符a,或b或c
	fmt.Println(b4)

	b5 := regexp.MustCompile(`[a-z]`).MatchString(`A23abc`) //必须是小写字母开头
	fmt.Println(b5)

	b6 := regexp.MustCompile(`[a-zA-Z][a-z]c`).MatchString(`Abcabc`) //字母
	fmt.Println(b6)

	b7 := regexp.MustCompile(`[0-9]`).MatchString(`110abc`) //数字
	fmt.Println(b7)

	b8 := regexp.MustCompile(`[^abc]`).MatchString(`*&%ABC123abc`) // 非abc开头即可
	fmt.Println(b8)

	b9 := regexp.MustCompile(`[[^a-z]`).MatchString(`memeda`) //
	fmt.Println(b9)

	b10 := regexp.MustCompile(`1[^2345]`).MatchString(`1a6`) //
	fmt.Println(b10)

	// 3.特殊符号
	b11 := regexp.MustCompile(`^\d`).MatchString(`a123`) //[0-9]
	//如果是""，\需要转义，``不需要转义
	fmt.Println(b11) //此处为true，是因为只要匹配到一个0-9的数字就可以，如果想以数字开头，可以在正则表达式前加^：`^\d`

	b12 := regexp.MustCompile(`\d\d`).MatchString(`1a2bc`) //[0-9]
	fmt.Println(b12)

	b13 := regexp.MustCompile(`\D`).MatchString(`*123abc`) //
	fmt.Println(b13)

	b14 := regexp.MustCompile(`^\w\w`).MatchString(`+-memeda`) //\w：[a-zA-Z0-9_] ，以\w\w开头
	fmt.Println(b14)

	b15 := regexp.MustCompile(`\d\w..`).MatchString(`123memeda`) //[0-9][a-zA-Z0-9_]任意字符任意字符
	fmt.Println(b15)

	b16 := regexp.MustCompile(`\W`).MatchString(`\nmemeda`) //
	fmt.Println(b16)

	b17 := regexp.MustCompile(`\w\W`).MatchString(`a*b = 20`) //
	fmt.Println(b17)

	b18 := regexp.MustCompile(`a[b-d]\w\d\W`).MatchString(`abc1\tdefg`) //
	fmt.Println(b18)
	b19 := regexp.MustCompile(`a\s`).MatchString(`a bc`) //匹配空白字符。即空格
	fmt.Println(b19)



}
