package main

import (
	"fmt"
)

func main() {

	//var a = []int{1, 2, 3, 4} //动态数组,切片slice,引用传递
	//fmt.Printf("a.types=%T\n", a)
	var b = make([]int, 4, 6)
	b[0] = 1
	////slice要声明容量，否则会报错
	//var c []int//runtime error: index out of range [0] with length 0
	//c[0] = 1
	//fmt.Println(b, a)

	//声明类型,初始元素，容量
	fmt.Printf("len=%d ,cap=%d,slice=%v\n", len(b), cap(b), b) //len=4 ,cap=6,slice=[1 0 0 0]

	//追加元素
	b = append(b, 1, 2)
	fmt.Printf("len=%d ,cap=%d,slice=%v\n", len(b), cap(b), b) //len=6 ,cap=6,slice=[1 0 0 0 1 2]

	//cap满之后再追加,cap翻倍
	b = append(b, 3)
	fmt.Printf("len=%d ,cap=%d,slice=%v\n", len(b), cap(b), b) //len=7 ,cap=12,slice=[1 0 0 0 1 2 3]

	//slice切片截取
	//slice[开始索引(包含),结束索引(不包含)] 左闭右开
	//但是只是len为4,容量cap还是b的容量大小,cap=12
	//并且，如果对截取的结果进行修改，原来的slice也会发生相应的改变(浅拷贝)
	c := b[0:4] //截取b下标为0到3的元素
	c[1] = 99
	fmt.Printf("len=%d ,cap=%d,slice=%v\n", len(c), cap(c), c) //len=4 ,cap=6,slice=[1 0 0 0]
	fmt.Printf("len=%d ,cap=%d,slice=%v\n", len(b), cap(b), b)

	d := make([]int, 4)
	copy(d, c)
	fmt.Printf("len=%d ,cap=%d,slice=%v\n", len(d), cap(d), d)
	//对d进行修改，查看c是否会发生变化(不会).因此copy是深拷贝(创建一个新的引用对象)
	d[2] = 100
	fmt.Printf("len=%d ,cap=%d,slice=%v\n", len(d), cap(d), d)
	fmt.Printf("len=%d ,cap=%d,slice=%v\n", len(c), cap(c), c)
}
