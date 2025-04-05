package main

import (
	"fmt"
	"strconv"
)

const x = 1

func main() {
	var number = 1
	fmt.Println(number)

	//Now instead of conventional initialisation of type, it can be inferred using the short assignment statement (:=)
	temp := 10
	var j int
	i := j
	mile, comp := 803, "moto"
	fmt.Println(comp)
	fmt.Println(mile)
	name := "hammad"
	fmt.Println(i + j)
	fmt.Println(strconv.Itoa(temp) + " " + name)
	//Only same data types, implying GO is strongly typed

	//GO follows the printf tradition from the C language

	s := fmt.Sprintf("I am %v years old", 22)
	fmt.Println(s)

}

/*
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128



Outside of a function (in the global/package scope), every statement begins with a keyword
( var, func, and so on) and so the := construct is not available.


fmt.Printf – Prints a formatted string to standard out
fmt.Sprintf() – Returns the formatted string
*/
