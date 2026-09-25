package main

import "fmt"

func main(){
	// pointers store the mem address of any val

	// &x -> address of X (makes a pointer )

	// *p -> dereference (go to that address and you can read or write it)

	score := 10
	fmt.Println("before:", score)

	addScore(&score)
	fmt.Println("after:", score)


}

func addScore(score *int){
	*score = *score + 5
}
