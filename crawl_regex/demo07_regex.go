package main

import (
	"regexp"
	"fmt"
)

func main() {

	// 边界问题

	b1 := regexp.MustCompile(`1[34578]\d{9}`).MatchString(`13278652345`)
	fmt.Println(b1) //true

	b2 := regexp.MustCompile(`1[34578]\d{9}`).MatchString(`1327865234578484848955`)
	fmt.Println(b2) //true

	b3 := regexp.MustCompile(`^1[34578]\d{9}$`).MatchString(`13278652345`) //
	fmt.Println(b3)                                                        //true

	b4 := regexp.MustCompile(`^[1-9]\d{17}$`).MatchString(`12345619881023432132`) //
	fmt.Println(b4)                                                               //false

	// 单词边界 "today is good"   (空格/开头)单词(末尾空格/结束)
	fmt.Println(regexp.MustCompile(`^\w+ve`).FindString(`hover`))          //hove
	fmt.Println(regexp.MustCompile(`\w+ve`).FindString(`hoverhoverhover`)) // hoverhoverhove
	fmt.Println(regexp.MustCompile(`^\w+ve$`).FindString(`hover`))         // ""
	fmt.Println(regexp.MustCompile(`^\w+ve`).FindString(`hover hover`))    //hove
	fmt.Println(regexp.MustCompile(`\w+ve\b`).FindString(`hoverhover`))    //""
	fmt.Println(regexp.MustCompile(`\w+ve\b`).FindString(`hove r`))        //hove
	fmt.Println(regexp.MustCompile(`^\w+\sve\b`).FindString(`ho ve r`))    //ho ve

}
