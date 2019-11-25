package main

import "fmt"

func main()  {
	nums :=[]int{78,109,2,563,300}
	largest(nums)
}

func largest(nums []int)  {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _,i :=range nums {
		if i > max {
			max = i
		}
	}
	fmt.Println("Largest number in", nums,"is",max)
}

func finished()  {
	fmt.Println("Finished finding largest")
}
//package main
//
//import "fmt"

//func main() {
//	slice := []int{0, 1, 2, 3}
//	myMap := make(map[int]*int)
//
//	for index, value := range slice {
//		num := value
//		myMap[index] = &num
//	}
//	fmt.Println(myMap)
//	fmt.Println("=====new map=====")
//	prtMap(myMap)
//}
//
//func prtMap(myMap map[int]*int) {
//	for key, value := range myMap {
//		fmt.Printf("map[%v]=%v\n", key, *value)
//	}
//}