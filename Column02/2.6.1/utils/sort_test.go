package utils

import (
	"fmt"
	"testing"
)

func TestSortStr(t *testing.T) {
	word := "zyaxwv"
	result := SortStr(word)
	fmt.Printf("result = %s\n", result)
	if result != "avwxyz" {
		t.Error("sort error.")
	}
}
