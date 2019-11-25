package main

import "fmt"

func main()  {
	a :=5
	b :=10
	var add,multiplied int
	add,multiplied = SumAndProduct(a,b)
	fmt.Println(add,multiplied)
}
func SumAndProduct(a,b int) (add int, mutiplied int)  {
	add = a+b
	mutiplied = a*b
	return
}
