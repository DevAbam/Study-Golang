package main
import (
	"errors"
	"fmt"
)

func main(){
	fmt.Println("Case 1: success")
	if err := doWork(true); err != nil{
		fmt.Println("error:", err)
	}

	fmt.Println("Case 2: fail early")
	if err := doWork(false); err != nil{
		fmt.Println("error:", err)
	}
}

func doWork(success bool) error {
	
	fmt.Println("start: resource acquired ")

	//defer will guarantee this runs at the end of this func
	// it will run in both paths (success and error return)
	defer fmt.Println("cleanup: resource released")

	if !success{
		return errors.New("something went wrong. im returning early")
	}
	fmt.Println("work: doing something important")
	fmt.Println("This work is done")

	return nil
}