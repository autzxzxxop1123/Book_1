package main 

import "fmt"

func makeeven() func() int {
	even := 0
	return func() int {
		even = even + 2 
		return even
	}
}
func main(){
	nexteven := makeeven()
	fmt.Println(nexteven())
	fmt.Println(nexteven())
	fmt.Println(nexteven())
}