package main

import "fmt"

type User struct{
	Name string
	Age int
}

func (u *User) Birthday(){
	u.Age++
}

func main(){
	u := User{Name: "Lord", Age: 13}
	fmt.Println(u.Age)

	u.Birthday()
	fmt.Println("After: ", u.Age)

}

