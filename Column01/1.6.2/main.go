package main

import (
	"fmt"
	"strconv"
)

func main() {
	i1, _ := strconv.ParseInt("010000", 2, 0)
	i2, _ := strconv.ParseInt("010101", 2, 0)
	fmt.Printf("i1 = %v\n", i1)
	fmt.Printf("i2 = %v\n", i2)

	r1 := i1 & i2
	fmt.Printf("r1 = %v\n", r1)

	r2 := i1 | i2
	fmt.Printf("r2 = %v\n", r2)

	i1 <<= 1
	fmt.Printf("i1 = %v\n", i1)

	i1 >>= 1
	fmt.Printf("i1 = %v\n", i1)
}
