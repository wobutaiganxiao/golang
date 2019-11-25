package main

import "fmt"

func main() {
	//1.bool
	// var a bool
	// a = true

	// var b bool = true

	// var c = true

	// d := true

	// var e bool  //缺省值为false
	// fmt.Println(a, b, c, d, e)

	//2.float32 float64
	// a := 3.14 //float64
	// fmt.Printf("a is type of %T\n", a)

	// var b float64 //缺省值为0
	// fmt.Println(b)

	//3.byte ACSII
	// var (
	// 	a byte = 97
	// 	b      = 65
	// 	c      = 48
	// )
	// fmt.Printf("%c, %d %c, %d, %c, %d\n", a, a, b, b, c, c)

	// a = 'a'
	// b = 'A'
	// c = '0'
	// fmt.Printf("%c, %d %c, %d, %c, %d\n", a, a, b, b, c, c)
	// fmt.Printf("%c\n", 'a'-32)
	// fmt.Printf("%d\n", '9'-'0')

	//4.string
	// var a string = "abcd"
	// fmt.Printf("%T, %s\n", a, a)

	// //求长度 len() 占内存的长度，不是字符个数
	// fmt.Printf("length:%d\n", len(a))
	// b := "那都不是事"
	// fmt.Printf("length:%d", len(b))
	//字符和字符串的关系
	//go内部的编码格式 —— utf-8
	//javaScript,JAVA unicode

	//解释型字符串和非解释型字符串、
	//非解释型字符串，不转译字符，不受换行影响
	str1 := "abcd\n"
	str2 := `abcd\n` //非解释型字符串
	fmt.Println(str1, str2)

	divDom := `
	<div>
		<div>
			<img/>
		</div>
		<p>
		</p>
	</div>`
	fmt.Println(divDom)
	fmt.Printf("%c", str1[2])
}
