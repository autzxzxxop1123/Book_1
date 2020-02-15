package main 

import "fmt"

func main(){

	number := 2
	switch number {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("Five")
	case 5:
		fmt.Println("six")
	default:
		fmt.Println("unknow")
		
	}
}