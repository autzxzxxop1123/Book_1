package main

import "fmt"

type books struct{
	title string
	auntor string
	subtitle string
	price float64
}
func main(){
Golang:=books{title:"Go Progarmming", price:200}
fmt.Println(Golang)
}