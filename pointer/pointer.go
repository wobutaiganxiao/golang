package main

import "fmt"

func main() {
	//1.定义int类型变量
	a := 10
	fmt.Println("a的数值是：", a)     //10
	fmt.Printf("%T\n", a)        //int
	fmt.Printf("a的地址是：%p\n", &a) //0xc00000a0b8
	//2.创建一个指针变量，用于储存a的地址
	var p1 *int
	fmt.Println(p1)//<nil>
	p1 = &a
	fmt.Println("p1的数值：", p1) //p1中存储的是a的地址
	fmt.Println("p1自己的地址:",&p1)
	fmt.Println("p1的数值是，是a的地址，该地址存储的数据： ",*p1)//获取指针指向的变量的数值
	//3.操作变量，更改数值并不会改变地址
	a = 100
	fmt.Println(a)
	fmt.Printf("%p\n",&a)
	//4.通过指针，改变变量的数值
	*p1 = 200
	fmt.Println(a)
	//5.指针的指针
	var p2 * *int
	fmt.Println(p2)
	p2 = &p1
	fmt.Printf("a: %T,  p1: %T,  p2: %T\n",a,p1,p2)//int,*int,**int(指针的指针)
	fmt.Println("p2的数值：",p2)//p1的地址
	fmt.Printf("p2自己的地址：%p\n",&p2)
	fmt.Println("p2中存储的地址，对应的数值，就是p1的地址： ",*p2)
	fmt.Println("p2中存储的地址，对应的数值，再获取对应的数值：",**p2)

}