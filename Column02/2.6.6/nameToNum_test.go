package main

import (
	"fmt"
	"testing"
)

func TestNameToNum(t *testing.T) {
	result := nameToNum("LESK", "M")
	fmt.Printf("result = %s\n", result)
	if result != "5375*6*" {
		t.Error("sort error.")
	}
}
