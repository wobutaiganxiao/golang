package main

import "fmt"

func max(n1, n2 int) int {
	var result int
	if n1 > n2 {
		result = n1
	} else {
		result = n2
	}
	return result
}
func swap(x, y, z string) (string, string, string) {
	return z, y, x
}

func main() {
	// a := 100
	// b := 200
	// var ret int
	// ret = max(a, b)
	// fmt.Printf("最大值：%d\n", ret)
	a, b, c := swap("Google", "Apple", "Micosoft")
	fmt.Println(a, b, c)
}
