package main

import "fmt"

//len -> how element you currently have
//capacity -> how many element you can store

func main() {
	//make ([], lem, cap) creates a slice with a given length and capacity
	scores := make([]int,0,5)

	fmt.Println(scores, len(scores), cap(scores))

	
	scores = append(scores, 100)
	fmt.Println("after appending 100", scores)
	scores = append(scores, 100, 200)
	fmt.Println("after appending 100, 200", scores)
	scores = append(scores, 45, 55)
	fmt.Println("after appending 45, 55", scores)
	
	//if i exceed the capaity, go grows the backing array (most at times doubles)
	scores = append(scores, 60)
	fmt.Println("after appending 60", scores, len(scores), cap(scores))

	todos := []string{"cook rice","wash dishes"}
	others := []string{"learn Golang"}

	// ... spread just like in JS
	todos = append(todos, others...)

}