package main

import "fmt"

func sum(number...int) int {
	total := 0 
	for _, n := range number {
		total = total + n
	}
	return total
}

func main(){
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9,}
	x := sum(number...)
	fmt.Println(x)
}