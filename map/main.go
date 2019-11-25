package main

import "fmt"

func main()  {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 10000
	fmt.Println("Original person salary", personSalary)
}