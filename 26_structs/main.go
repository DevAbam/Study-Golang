package main

import "fmt"

type User struct{
	ID int
	Name string
	Email string
	Age int
}

func main(){
	u1 := User{ID: 1, Name: "Ramesh,", Email: "Ramesh@gmail.com", Age: 29}

	fmt.Println(u1)
}