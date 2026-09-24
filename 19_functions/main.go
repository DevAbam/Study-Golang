package main

import "fmt"

func add(a int , b int) int {
	return a + b ;
}

func sumAndProduct(a int, b int) (int, int){
	sum := a + b
	product := a*b
	return sum, product
}

func main(){

	sum1 := add(10, 20)
	s , p := sumAndProduct(3,3)

	onlySum , _ := sumAndProduct(10,2)

	fmt.Println(onlySum, s, p)
	fmt.Println(sum1)

}