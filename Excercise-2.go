package main

import "fmt"

func add(val int) (x, y int) {

	x = val + 4*10
	y = val / 10
	return

}

func main() {

	fmt.Print(add(10))
}
