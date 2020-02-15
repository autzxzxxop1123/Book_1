package main

import "fmt"

type books struct{
	title string
	auntor string
	subtitle string
	price float64
}
func main(){
var book1 books
book1.title=" Go Programming"
book1.auntor="GO"
book1.subtitle="Go tutorial"
book1.price=200
fmt.Println(book1.price)
fmt.Println(book1.subtitle)
fmt.Println(book1.auntor)
fmt.Println(book1.title)
}