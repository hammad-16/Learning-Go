package main

import (
	"fmt"
)

func display(arr []int) []int {
	return arr
}
func main() {

	num := 5
	ptr := &num
	fmt.Println(ptr, "", num)

	var arr [2]int

	arr[0] = 1
	arr[1] = 2
	fmt.Println(display(arr[:])) // Here you are converting an array into slice

	var slice []int
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)
	slice = append(slice, 4, 5)
	fmt.Println(slice)
	//eg . myArray := [5]int{1, 2, 3, 4, 5}

	//A slice is like a dynamic array in Golang that can grow or shrink as needed. It’s basically a view into an underlying array.
}

/*be careful when working with pointers because they can easily
lead to null pointer errors and memory leaks if not used correctly.
Arrays: [n]int
Arrays always declared with fixed sizes

Empty Slice: var s []int
s = append(s,10,20)


*/
