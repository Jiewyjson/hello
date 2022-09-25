package main

import (
	"fmt"
)

func test1(a int, b string) (r1 string, r2 int) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("r1=", r1)
	fmt.Println("r2=", r2)

	r1 = "go"
	r2 = 100
	return
}

func main() {
	//说明string的默认值为空
	println("Hello, go!")
	var x, y = test1(10, "helloGo")
	println("x=", x)
	println("y=", y)
}
