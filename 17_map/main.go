package main

import "fmt"

func main() {
	//map[keytype]valuetype

	ages := map[string]int{
		"abam": 65,
		"john": 35,
	}

	fmt.Println(ages["abam"], len(ages))

	//make(map[k]v)

	scores := make(map[string]int)

	scores["math"] = 90

	users := map[string]string{
		"u1": "ramesh",
		"u2": "john",
		"u3": "jane",
	}

	delete(users, "u2")
	fmt.Println(users)
}
