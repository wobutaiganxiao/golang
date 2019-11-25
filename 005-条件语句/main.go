package main

import (
	"fmt"
)

func main() {
	// if
	// 1.if 没有括号
	// n := 5
	// if n > 0 {
	// 	fmt.Println("Great!")
	// }

	// if n < 0 {
	// 	fmt.Println("Path 1")
	// } else {
	// 	fmt.Println("Path 2")
	// }
	//3.特殊写法
	//	第一，紧凑
	//	第二，contents,err仅仅在作用域内
	// 	const filename = "1.txt"
	// 	contents, err := ioutil.ReadFile(filename)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Printf("%s\n", contents)
	// 	}
	// const filename = "1.txt"
	// if contents, err := ioutil.ReadFile(filename); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%s\n", contents)
	// }
	// 	fmt.Printf("%s", contents)
	//4.switch 没有括号,不需要写break
	// a := "hello"
	// switch a {
	// case "hello":
	// 	fmt.Println("1")
	// case "world":
	// 	fmt.Println("2")
	// default:
	// 	fmt.Println("3")
	// }
	//5.一分支多值
	// a := "hello"
	// switch a {
	// case "hello", "Hello":
	// 	fmt.Println("1")
	// case "world":
	// 	fmt.Println("2")
	// default:
	// 	fmt.Println("3")
	// }
	//6.分支表达式,case首先匹配为true的，运行
	//7.fallthrough,用于兼容C\C++
	//成绩
	socre := 95

	switch {
	case socre > 100 || socre < 0:
		//fmt.Println("Error")
		panic("Error")
	case socre >= 90:
		fmt.Println("A")
		//fallthrough
	case socre >= 70:
		fmt.Println("B")
	case socre >= 60:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}

	// 特点回顾：
	//1. if, switch没有括号
	//2. if里可以定义变量
	//3. switch不需要break
}
