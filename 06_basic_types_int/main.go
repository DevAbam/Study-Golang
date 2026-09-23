package main
import (
	"fmt"
)

func main(){
	views1 := 1000
	views2 := 2000
	totalView := views1 + views2 

	likes := 10
	likes++

	avgViews := totalView/2

	fmt.Println(totalView, likes, avgViews)
}