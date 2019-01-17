package main

import (
	"regexp"
	"fmt"
)

func main() {

	// 分组

	b1 := regexp.MustCompile(`[0-9]|x`).MatchString(`x`)
	fmt.Println(b1)

	b2 := regexp.MustCompile(`[ab]`).MatchString(`a`)
	fmt.Println(b2)

	b3 := regexp.MustCompile(`abc|ddd`).MatchString(`ddd`)
	fmt.Println(b3)

	// 身份证：  练习3：身份证号：18位。0不能开头第一位：非0,16位。最后一位：(数字|X)
	fmt.Println(regexp.MustCompile(`[1-9]\d{17}|[1-9]\d{16}X`).MatchString(`13141219880712321X`))
	fmt.Println(regexp.MustCompile(`^[1-9]\d{16}(\d|X)$`).MatchString(`13141219880712321X`))

	//练习2：邮箱：163.com，qq.com?
	s1 := `wangergou@163.com`
	s2 := `sanpang@qq.com` // 828384848@qq.com
	s3 := `lixiaohua@sina.com`

	fmt.Println(regexp.MustCompile(`\w+@(163|qq)\.com`).MatchString(s1))
	fmt.Println(regexp.MustCompile(`\w+@(163|qq)\.com`).MatchString(s2))
	fmt.Println(regexp.MustCompile(`\w+@(163|qq)\.com`).MatchString(s3))

	// 练习3：邮箱：163.com,sina.cn,yahoo,cn,qq.com
	fmt.Println(regexp.MustCompile(`\w+@(163|sina|yahoo|qq)\.(com|cn)`).MatchString(`sanpang@163.com`))

	fmt.Println("----------------------")

	s4 := `<html><h1>helloworld</h1></html>李小花李小花`
	re1 := regexp.MustCompile(`<(.+)><.+>(.+)</.+></.+>`)

	//打印分组的数量
	fmt.Println(re1.NumSubexp()) //2

	res1 := re1.FindAllStringSubmatch(s4, -1)
	fmt.Println(res1)
	fmt.Println(res1[0])
	fmt.Println(res1[0][0])
	fmt.Println(res1[0][1])
	fmt.Println(res1[0][2])

	fmt.Println("-----------------------")

	s5 := `<html><body><h1>hello</h1></body></html>`
	re2 := regexp.MustCompile(`<(?P<t1>.+)><(?P<t2>.+)><(?P<t3>.+)>(?P<t4>.+)</(.+)></(.+)></(.+)>`)

	//获取分组名称
	for i := 0; i <= re2.NumSubexp(); i++ {
		fmt.Printf("%d: %q\n", i, re2.SubexpNames()[i])
	}

	res2 := re2.FindAllStringSubmatch(s5, -1)
	fmt.Println(res2)

}
