package main

import "fmt"

func main() {

	var x int = 5
	var y int = 30
	var z int = x * y

	fmt.Println("cevap")

	if x > -6 && y >= 30 || x > y && z%x > y+z {

		fmt.Println("dogru")
	} else {

		fmt.Println("yanlis")
	}
}
