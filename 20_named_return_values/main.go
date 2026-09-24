package main

import "fmt"

func prodAndSum(num1 int, num2 int)( prod int , sum int){
	prod = num1*num2
	sum = num1 + num2
	return  // naked return 
	
}

func main(){
	p , s := prodAndSum(10,10)
	fmt.Println(p, s)
}