package main

import (
	"regexp"
	"fmt"
)

func main() {

	// 数量词

	b1 := regexp.MustCompile(`\d*`).MatchString(`123`)
	fmt.Println(b1)

	b2 := regexp.MustCompile(`\d*`).MatchString(`abc123`)
	fmt.Println(b2)

	b3 := regexp.MustCompile(`\d*`).MatchString(``)
	fmt.Println(b3)

	b4 := regexp.MustCompile(`a[b-f]*\d\d[xy]*`).MatchString(`abbbb123x`) //至少0位，意味着可以没有
	fmt.Println(b4)

	b5 := regexp.MustCompile(`a[b-f]*\d\d[xy]*`).MatchString(`a12xxyxx3x`)
	fmt.Println(b5)

	b6 := regexp.MustCompile(`\d+abc`).MatchString(`abcmemeda`) //+至少一位
	fmt.Println(b6)

	b7 := regexp.MustCompile(`\d+.*\w+`).MatchString(`1234\ncd`) //
	fmt.Println(b7)

	b8 := regexp.MustCompile(`\d?[a-z]+`).MatchString(`123abc`) //
	fmt.Println(b8)

	b9 := regexp.MustCompile(`\d?\w+`).MatchString(`123abc`) //
	fmt.Println(b9)

	b10 := regexp.MustCompile(`\d{4}[a-z]+`).MatchString(`1234abcd`) //\d刚好4次
	fmt.Println(b10)

	b11 := regexp.MustCompile(`\d{4,}[a-z]+`).MatchString(`12345abcd`) //\d至少4次
	fmt.Println(b11)

	// +-->1次或多次，至少1次{1,}

	b12 := regexp.MustCompile(`\d{4,6}[a-z]+`).MatchString(`1234567abcd`) //
	fmt.Println(b12)

	// ?-->0次或1次，{0,1}

	//1.匹配手机号码：
	/*
	13012345678, 131xxxxxxxx,132xxxxxxxx,133xxxxxxxx,134xxxxxxxx,135xxxxxxxx,136xxxxxxxx,137xxxxxxxx,138,139
	# 第一位：1，第二位：34578，第三位：0-9   11位。
	*/
	b13 := regexp.MustCompile(`1[34578]\d{9}`).MatchString(`13212344321`)
	fmt.Println(b13)

	// 2.验证QQ号：第一位非0，长度：5位-11位。
	//441883704

	b14 := regexp.MustCompile(`[1-9]\d{4,10}`).MatchString(`44188370445`)
	fmt.Println(b14)

	/*
	练习1：日期：2017-11-29
		年份是4位数字，月份是1-2位数字，日期1-2位数字
	练习2：邮箱：163.com，qq.com?

	*/

	b15 := regexp.MustCompile(`\d{4}-\d{1,2}-\d{1,2}`).MatchString(`2018-12-25`)
	fmt.Println(b15)

	b16 := regexp.MustCompile(`[1-9]\d{17}`).MatchString(`231412198807123214`)
	fmt.Println(b16)

}
