package main 

import "fmt"

func main(){
	fmt.Print("Enter your Number:")
	var number int
	fmt.Scanf("%d", &number)
	switch number {
	case 0: fmt.Println("zero")
	case 1: fmt.Println("one")
		
	}
}