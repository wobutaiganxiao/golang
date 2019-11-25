package main

import "fmt"

func a()  {
	//defer执行顺序为先进后出
	i :=0
	defer fmt.Println(i)
	i++
	defer fmt.Println(i)
	return
}
func main()  {
	a()
}