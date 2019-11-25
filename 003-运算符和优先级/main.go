package main

import "fmt"

func main() {
	//1.与&& 或|| 非！
	//与或非的优先级不一样，非最高，与其次，或最低
	//fmt.Println(true || false && true || !false)
	// true || false && ture ||
	// true || false ||
	// true || !false
	// true || true
	//true

	//2. == !==
	// > >= < <=
	// s1, s2, s3 := "10", "10", "9" //字符
	// fmt.Println(s1 == s2)
	// fmt.Println(s1 != s3)
	// fmt.Println(s1 > s3) //因为字符串“10”中的“1”对应ASCII中49，字符“9”对应57，所以字符串“10” < "9"
	// n1, n2, n3 := 10, 10, 9
	// fmt.Println(n1 == n2)
	// fmt.Println(n1 != n3)
	// fmt.Println(n1 > n3)

	//3. | 双目运算符，按位或
	//   ^ 双目运算符，按位异或:如果两个数字不一样，true, 一样，返回false
	// n1, n2 := 5, 11
	// fmt.Println(n1 | n2)
	// fmt.Println(n1 ^ n2)
	// //	^单目运算符，取反
	// //  如何一个负数的值 1.确定符号，2.取反加1
	// fmt.Println(^n1)
	//4. << 双目运算符 左移
	//	 >> 双目运算符 右移
	// n1 := 2
	// fmt.Println(4 << n1)
	// fmt.Println(4 >> n1)

	//5.& 双目运算符，按位与，A & mask，通过配置mask, 置零的位是不关心的位，置1的位是需要的位
	//	&^双目运算符，按位消除， A &^ mask,通过配置mask,置一的位抹成零，置零的位不变
	// n1 := 5
	// fmt.Println(n1 & 4)
	// fmt.Println(n1 &^ 4)
	//6. go 之简洁
	//   go中没有三目运算符
	//	 ++,-- 是语句，不是运算符，++ -- 是不能出现在表达式里面

	// n := 5
	// n++
	// b := n //b := n++是错误的
	// fmt.Println(n, b)

	// + - * / %
	// / %
	// n1, n2 := 13, 7
	// fmt.Println(n1 / n2)
	// fmt.Println(n1 % n2)

	//8. go语言中的预算符优先级
	// * / % << >> & &^
	// + - | ^
	// == != < <= > >=
	// &&
	// ||
	// fmt.Println(5*2 + 12%5<<2 | 5)
	//5*2 +
	//10+12%5<<
	//10+2<<2 |
	//10+8|5
	//18|5
	fmt.Println(5 > 4 || 3 < 2 && 7 > 8 || 4 <= 9)
	//5>4||
	//true || 3<2 &&
	//true || false && 7>8 ||
	//true || false && false ||
	//true || false ||
	//true || 4 <= 9
	//true || true
	// true
}
