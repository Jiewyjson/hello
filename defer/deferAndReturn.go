package main

import (
	"fmt"
)

func returnfunc() int {
	fmt.Println("returnfunc init..")
	return 0
}

func deferfunc() int {
	fmt.Println("deferfunc init..")
	return 1
}
func deferAndReturn() int {
	defer deferfunc()
	return returnfunc()
}

func main() {
	fmt.Println("main init..")
	deferAndReturn()
}
