package main 		

import "fmt"

func main(){
	for i := 0; i <=10; i = i+1 {
		if i%3 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd") 
		}
	}
}