package main

import "fmt"

func main() {
	a := 1
	b := 2
	sum(a, b)

	a = 10
	b = 9
	sum(a, b)

	a = 39
	b = 10
	sum(a, b)

	a = 5
	b = 2
	sum(a, b)

	a = 2
	b = 2
	sum(a, b)

	a = 6
	b = 3
	sum(a, b)

	a = 50
	b = 100
	sum(a, b)

	a = 60
	b = 40
	sum(a, b)

	a = 15
	b = 12
	sum(a, b)

	a = 100
	b = 99
	sum(a, b)
}

func sum(a int, b int) {
	c := a*3 - b
	fmt.Println(a, "* 3 -", b, "=", c, "!")
}
