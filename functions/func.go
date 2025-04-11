package main

import "fmt"


func add(x int, y int) int { //function signature

	return x - y
}

func main() {
	fmt.Println(add(5, 4))
}
