package main

import "fmt"

func main() {
//	slice_1 := []int{1, 2, 3, 4, 5}
//	fmt.Printf("main-->data:\t%#v\n", slice_1)
//	fmt.Printf("main-->len:\t%#v\n", len(slice_1))
//	fmt.Printf("main-->cap:\t%#v\n", cap(slice_1))
//	test1(slice_1)
//	fmt.Printf("main-->data:\t%#v\n", slice_1)
//
//	test2(&slice_1)
//	fmt.Printf("main-->data:\t%#v\n", slice_1)
//}
//func test1(slice_2 []int) {
//	slice_2[1] = 6666               // 函数外的slice确实有被修改
//	slice_2 = append(slice_2, 8888) // 函数外的不变
//	fmt.Printf("test1-->data:\t%#v\n", slice_2)
//	fmt.Printf("test1-->len:\t%#v\n", len(slice_2))
//	fmt.Printf("test1-->cap:\t%#v\n", cap(slice_2))
//}
//
//func test2(slice_2 *[]int) { // 这样才能修改函数外的slice
//	*slice_2 = append(*slice_2, 6666)
//
//	numbers := make([]int,5)
//	fmt.Println(reflect.TypeOf(numbers))
//	var balance = [5]int{1000, 2, 3, 17, 50}
//	copy(numbers,balance)
//	fmt.Println(reflect.TypeOf(numbers))
//	var numbers []int
//	printSlice(numbers)
//	numbers = append(numbers,0)
//	printSlice(numbers)
//	numbers = append(numbers,1)
//	printSlice(numbers)
//	numbers = append(numbers,2,3,4)
//	printSlice(numbers)
//	numbers1 := make([]int, len(numbers),(cap(numbers)*2))
//	printSlice(numbers1)
//	copy(numbers1,numbers)
//	printSlice(numbers1)
//	numbers1[0] = 1
//	printSlice(numbers)
//	printSlice(numbers1)
	a :=[]int{0,1,2,3,4}
	b :=a
	fmt.Println(a)
	b[2]=1
	fmt.Println(a)
	fmt.Println(b)
	printSlice(a)
	a = append(a,5)
	printSlice(a)
}
func printSlice(x []int)  {
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}