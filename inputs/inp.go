package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var x int
	fmt.Scanln(&x)
	fmt.Println("value of : ", x)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter name")
	tex, err := reader.ReadString('\n')

	//This returns a string and an error
	//If we use _ in place of err it means that we are discarding the second return value
	fmt.Println(tex)
	fmt.Println(err)
	var start, end int
	fmt.Scan(&start)
	fmt.Scan(&end)

	for i := start; i <= end; i++ {

		fmt.Println(i, " ", i+1)

	}

}

/* Scan stops at whitespace.

Scanln stops at newline.

Scanf allows formatted input (like scanf in C).
*/
