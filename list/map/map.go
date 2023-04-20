package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func main() {
	m := make(map[int]Product)

	m[1] = Product{"볼펜", 500}
	m[2] = Product{"지우개", 200}
	m[3] = Product{"자", 1000}
	m[10] = Product{"샤프", 3500}
	m[111] = Product{"샤프심", 1500}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
