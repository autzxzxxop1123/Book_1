package main

import "fmt"

func say(){
	fmt.Println("hello")
}

func greet(name string){
	fmt.Println("hello", name)
}

func main(){
	say()
	greet("world")
}