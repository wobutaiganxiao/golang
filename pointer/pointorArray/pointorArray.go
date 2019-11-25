package main

import "fmt"

func main()  {
	//数组指针：首先是一个指针，一个数组的地址
	//1.创建一个普通是数组
	arr1 :=[4]int{1,2,3,4}
	fmt.Println(arr1)

	//2.创建一个指针，存储该数组的地址--->数组指针
	var p1 *[4]int
	p1 = &arr1
	fmt.Println(p1)//&[1 2 3 4]
	fmt.Printf("p1中所存储的值:%p\n",p1)//数组arr1的地址
	fmt.Printf("p1自己的地址:%p\n",&p1)

}
