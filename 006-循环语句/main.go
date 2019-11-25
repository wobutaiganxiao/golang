package main

import (
	"fmt"
)

func convertTOBin(n int) string {
	if n < 0 {
		return ""
	}
	result := ""

	for ; n > 0; n /= 2 {
		latBin := n % 2
		result = fmt.Sprintf("%d", latBin) + result
	}
	return result
}
func main() {
	//1. for 没有括号
	//	 for <初始条件>；<循环条件>;<循环后条件>{}
	// sum := 0
	// for i := 1; i <= 100; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	//2.初始条件可以省略

	//查找字符串中第一个字符'c'的位置
	// str := "abcde"
	// i := 0
	// for ; i < len(str); i++ {
	// 	if str[i] == 'c' {
	// 		break
	// 	}
	// }
	// fmt.Printf("i = %d\n", i)
	//3.初始条件和循环条件都可以省略
	// str := "abcde"
	// i := 0
	// for ; ; i++ {
	// 	if i >= len(str) || str[i] == 'c' {
	// 		break
	// 	}
	// }
	// fmt.Printf("i = %d\n", i)
	//4.全部省略
	// str := "abcde"
	// i := 0
	// for {
	// 	if i >= len(str) || str[i] == 'c' {
	// 		break
	// 	}
	// 	i++
	// }
	// fmt.Printf("i = %d\n", i)

	//5.只有循环条件
	// str := "abcde"
	// i := 0
	// for i < len(str) {
	// 	if str[i] == 'c' {
	// 		break
	// 	}
	// 	i++
	// }
	// fmt.Printf("i = %d\n", i)
	//6.go语言中没有while,do-whlie
	//读文件
	// if file, err := os.Open("./1.txt"); err != nil {
	// 	panic(err)
	// } else {
	// 	scanner := bufio.NewScanner(file)
	// 	for scanner.Scan() {
	// 		fmt.Println(scanner.Text())
	// 	}
	// }
	fmt.Println(convertTOBin(15))
}
