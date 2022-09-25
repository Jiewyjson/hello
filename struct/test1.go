package main

import "fmt"

// 定义一个type:book
type book struct {
	title  string
	author string
}

func main() {
	//初始化一个book
	var b book
	b.title = "Go语言编程"
	b.author = "jiege"
	println(b.title, b.author)
	fmt.Println("--------")
	changeBook(b)

}
func changeBook(book book) {
	book.author = "jiege2"
	fmt.Println(book.title, book.author)
}
