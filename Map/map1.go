package main

import "fmt"

func main() {
	//定义一个map
	map1 := make(map[string]string)
	//添加元素
	map1["city1"] = "guangzhou"
	map1["city2"] = "shenzhen"
	map1["city3"] = "beijing"

	//删除元素
	delete(map1, "city1")

	//修改元素
	map1["city2"] = "shanghai"
	chaneMap(map1)

	//遍历map
	for key, value := range map1 {
		fmt.Printf("key=%v,value=%v\n", key, value)
	}
	fmt.Println("---------")
	printMap(map1)
}

func printMap(map1 map[string]string) {
	for key, value := range map1 {
		fmt.Printf("key=%v,value=%v\n", key, value)
	}
}

func chaneMap(map1 map[string]string) {
	map1["city1"] = "guangzhou"
}
