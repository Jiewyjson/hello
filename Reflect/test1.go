package main

import (
	"fmt"
	"reflect"
)

type User struct {
	id   int
	Name string
	age  int
}

func (this User) Call() {
	fmt.Println("user is called..")
	fmt.Printf("this:%v", this)
}

func main() {
	user := User{1, "jiege", 22}
	getFileAndMethod(user)
}
func getFileAndMethod(input interface{}) {
	//获取函数名
	inputType := reflect.TypeOf(input)
	fmt.Printf("inputType:%v,", inputType.Name())
	//获取value
	inputValue := reflect.ValueOf(input)
	fmt.Printf("inputValue:%v,", inputValue)
	//获取字段
	//for i := 0; i < inputType.NumField(); i++ {
	//	field := inputType.Field(i)
	//	value := inputValue.Field(i).Interface()
	//	fmt.Printf("%s:%v", field.Name, value)
	//}
	//获取方法
	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)
		fmt.Printf("%s:%v", method.Name, method.Type)
	}
}
