package main

import (
	"regexp"
	"fmt"
)

func main() {

	//转义字符
	b1 := regexp.MustCompile(`\n\w`).MatchString(`\nabc`)
	fmt.Println(b1) // false

	b2 := regexp.MustCompile(`\n\w`).MatchString(`\\nabc`)
	fmt.Println(b2)//false

	b3 := regexp.MustCompile(`\\n\w`).MatchString(`\\nabc`)
	fmt.Println(b3)//true

	b4 := regexp.MustCompile(`\\\\n\w`).MatchString(`\\nabc`)
	fmt.Println(b4)

	b5 := regexp.MustCompile(`\\n\w`).MatchString(`\\nabc`)
	fmt.Println(b5)
	b6 := regexp.MustCompile(`\.\w+`).MatchString(`.abc`)
	fmt.Println(b6)


}
