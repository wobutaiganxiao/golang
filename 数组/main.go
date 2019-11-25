package main

import "fmt"

func main() {
	//格式：var variable_name [SIZE] variable_type
	// var n [10]int
	// for i := 0; i < 10; i++ {
	// 	n[i] = i + 1000
	// }
	// for j := 0; j < 10; j++ {
	// 	fmt.Printf("Element[%d] = %d\n", j, n[j])
	// }
	//二维数组
	// var arr = [5][2]int{
	// 	{0, 0},
	// 	{1, 2},
	// 	{2, 4},
	// 	{3, 6},
	// 	{4, 8},
	// }
	// for i := 0; i < 5; i++ {
	// 	for j := 0; j < 2; j++ {
	// 		fmt.Printf("arr[%d][%d] = %d\n", i, j, arr[i][j])
	// 	}
	// }
	var balance = [5]int{1000, 2, 3, 17, 50}
	var avg float32
	//数组作为参数传递给函数
	avg = getAverage(balance,5)
	fmt.Printf("平均值为：%f", avg)
}
func getAverage(arr [5]int, size int) float32 {
	var sum int
	var avg float32
	for i := 0; i < size; i++ {
		sum += arr[i]
	}
	avg = float32(sum) / float32(size)
	return avg
}
