package main

import "fmt"

func main() {

	adams := 42

	fmt.Printf("42 as binary, %b \n", adams)
	fmt.Printf("42 as hexadecimel, %x \n", adams)

	a, b, c, d, e, f := 0, 1, 2, 3, 7, 8

	fmt.Printf("Hexadecimel = %x is a, %x is b, %x is c, %x is d, %x is e, %x is f", a, b, c, d, e, f)
	fmt.Printf("Binary - %b is a, %b is b, %b is c, %b is d, %b is e, %b is f", a, b, c, d, e, f)
}
