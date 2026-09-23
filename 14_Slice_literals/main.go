package main

import "fmt"

//dynamic and can grow

func main() {
	res := []string{"Lord","Ramesh"}

	res[0] = "Abam"

	var nums []int 
	nums = append(nums, 10)
	nums = append(nums, 20,30)

	fmt.Println(res, res[0], res[len(res)-1])
	fmt.Println(nums)
}