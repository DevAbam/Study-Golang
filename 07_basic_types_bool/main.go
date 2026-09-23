package main
import (
	"fmt"
)

func main(){
	isLogged := true
	isAdmin := false
	hasSubsription := true

	canOpenDashbard := isLogged && hasSubsription

	canDeletePost := isAdmin || (isLogged && hasSubsription)

	fmt.Println(canOpenDashbard, canDeletePost)
}