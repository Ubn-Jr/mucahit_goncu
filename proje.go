package main

import "fmt"

func main() {

	var x int
	var y int
	var z int
	var m int
	var g int
	x = -20
	y = -5
	z = add(x, y)
	m = ert(z, x)
	g = we(m, y)
	fmt.Println(z, m, g)
}

func add(x int, y int) int {
	var z int
	z = x * y
	return z
}
func ert(z int, x int) int {
	var m int
	m = z - x
	return m
}
func we(m int, y int) int {
	var g int
	g = m % y
	return g
}
