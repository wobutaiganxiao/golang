package main

import (
	"fmt"
)

func main()  {
	a := 100
	b := 200
	var res int
	res =max(a,b)
	fmt.Println(res)
}

func max(num1 int, num2 int) int {
	if (num1 > num2) {
		return num1
	} else {
		return num2
	}
}