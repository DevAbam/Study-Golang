package main

import (
	"fmt" 
	"strings"
)

func main(){
	firstname := "Abam"
	lastname := "Adjetey"
	fullname := firstname + " " + lastname
	
	fmt.Println(fullname)

	fmt.Println(strings.ToUpper(fullname))
}