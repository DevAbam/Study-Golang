package main

import "fmt"

func main(){

	day := 3

	switch day {
	case 1:
		fmt.Println("Mon")
	case 2:
		fmt.Println("Tue")
	case 3:
		fmt.Println("wed")
	default:
		fmt.Println("Not a day")
	}
}