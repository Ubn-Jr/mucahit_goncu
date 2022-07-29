package main

import (
	"fmt"
)

func main() {
	var (
		a string
		b int
		c int
		d int
		e int
		f int
	)
	a = "mucahit"
	e = 5
	c = mg(b, e)
	f = 2
	d = gm(e, f)

	fmt.Println(a, b, c, d)

}

func mg(b int, e int) int {
	var c int
	c = b * e
	return c
}
func gm(e int, f int) int {
	var d int
	d = e + f
	return d

}
