package main

import (
	"fmt"
	"math"
	"math/big"
)

func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b
}

func main() {
	// 작은 실수 오차 무시하기
	var a = 0.1
	var b = 0.2
	var c = 0.3

	fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a+b)
	fmt.Printf("%0.18f + %0.18f = %v\n", c, a+b, equal(a+b, c))

	a = 0.000000000004
	b = 0.000000000002
	c = 0.000000000007

	fmt.Printf("%g == %g : %v\n", c, a+b, equal(a+b, c))

	// math/big 패키지 사용
	x, _ := new(big.Float).SetString("0.1")
	y, _ := new(big.Float).SetString("0.2")
	z, _ := new(big.Float).SetString("0.3")

	k := new(big.Float).Add(x, y)
	fmt.Println(x, y, z, k)
	fmt.Println(z.Cmp(k))
}
