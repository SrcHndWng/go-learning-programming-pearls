package main

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{16, 20, 69, 70, 78, 79, 82, 91, 97, 113}

	test := func(target int, index int, value int) {
		i, v := binarySearch(target, nums)
		fmt.Printf("target = %v, index = %v, value = %v\n", target, i, v)
		if i != index || v != value {
			t.Error("search error.")
		}
	}

	test(91, 7, 91)
	test(16, 0, 16)
	test(113, 9, 113)
	test(10, -1, -1)
	test(200, -1, -1)
	test(17, -1, -1)
	test(96, -1, -1)
}
