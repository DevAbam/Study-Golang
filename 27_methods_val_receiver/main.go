package main

import "fmt"


type User struct{
	Name string
	Age int
}
//functions attached to the struct.. (currently i tink of it as methods of a class)
func (u User) Intro() string{
	return fmt.Sprintf("Hi, i am %s", u.Name)
}

func main(){
	u := User{Name: "Sam", Age: 20}
	fmt.Println(u.Intro())
}