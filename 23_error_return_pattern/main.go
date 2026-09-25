package main

import (
	"fmt"
	"log"
	"strconv"
)

func main(){

	// go dnt use exceptions for normal failures
	// functions return errors as normal return values
	// val , err := something()
	//if err != nil {handle the error}
	if err := run(); err != nil{
		log.Fatal(err)
	}

}

func run() error {
	input := "30"
	level , err := parseLevel(input)
	if err != nil{
		return  err
	}
	fmt.Println("selectd level", level)
	return nil
}

func parseLevel(s string) (int , error){
	// (val , err)
	//(nil error -> success) (not nil -> failure)
	n , err  := strconv.Atoi(s)
	if err != nil{
		return 0, fmt.Errorf("Level must be a number")
	}
	if n < 1 || n > 5{
		return 0, fmt.Errorf("Level must be tetween 1 and 5")
	}
	return n , nil
}