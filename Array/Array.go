package main

import (
	"fmt"
	"reflect"
)

func main() {
	//创建一个定长数组，值传递
	var array1 [10]int
	array2 := [10]int{1, 3, 5, 7}

	//遍历的方法
	for i := 0; i < len(array1); i++ {
		println(i, "=", array1[i])
	}

	for index, value := range array2 {
		println("index=", index, "value=", value)
	}

	//查看数组类型
	fmt.Println("array1 types=", reflect.TypeOf(array1))
	fmt.Println("array2 types=", reflect.TypeOf(array2))

	printArray(array2)
}

func printArray(array [10]int) {
	for index, value := range array {
		println("index=", index, "value=", value)
	}
}
