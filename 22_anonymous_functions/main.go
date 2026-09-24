package main

import "fmt"

func main(){
	res := func(n int)int{
		return n * 2
	}

	fmt.Println(res(2))

	//IIFE  IMMEDIATELY INVOKED FUNCTION EXPRESSION

	res1 := func(a int, b int) int{
		return a + b
	}(2,3)

	fmt.Println(res1)
}
