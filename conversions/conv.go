package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello again")
	reader := bufio.NewReader(os.Stdin)

	inp, _ := reader.ReadString('\n')

	fmt.Println("String:", inp)
	//converting input into a number, and also checking if there is any error.
	//by default the input has a endline suffix, so to remove that we use strings.TrimSpace()
	//also if input is not a number, then we handle the error

	num, err := strconv.ParseInt(strings.TrimSpace(inp), 10, 32)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Number is ", num)
	}
}

/*
Other conversions:
strconv.Atoi
strconv.ParseInt
strconv.ParseFloat
strconv.Itoa
strconv.FormatInt
strconv.FormatUint
strconv.FormatFloat



panic is a built-in function in Go that immediately stops normal execution of the current function and begins panicking, which means:
It unwinds the call stack.
It runs any defer statements on the way up.
If not recovered, it crashes the program.


*/
