package main 

import "fmt"

func main(){
	var P1 *int
	fmt.Println("P1 :", P1)
	P2 := new(int)
	fmt.Println("P2 :", P2)
}