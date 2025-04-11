package main

import "fmt"

const x = 2

func main() {
	c := 2
	if c > 18 {
		fmt.Println("Virat")
	} else if c == 7 {
		fmt.Println("Dhoni")
	} else {
		fmt.Println("Goa")
	}
	fmt.Println("Anyways", x)

	var i int
	fmt.Scan(i)

	switch i {
	case 1:
		fmt.Println("Cristiano Ronaldo")
	case 2:
		fmt.Println("Messi")
	case 3:
		fmt.Println("Lol")
	default:
		fmt.Println("None")
	}
	//The conditional statements need to follow proper indentation and no paranthesis required

}
