package main

import "fmt"

func square(x int) int {
	return x * x
}

func main() {
	fmt.Printf("9 * 9 = %d\n", square(9))
}
