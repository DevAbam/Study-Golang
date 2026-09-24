package main

import "fmt"

func sumAll(nums ...int) int{
	total := 0
	for _ , num := range nums{
		total += num
	}
	return total
}

func main(){
	values := []int{10,20,30,40}
	fmt.Println(sumAll(values...))
	fmt.Println(sumAll(1,2,3,4))

}
