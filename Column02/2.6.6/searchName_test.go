package main

import (
	"fmt"
	"testing"
)

func TestSearchName(t *testing.T) {
	var names personNames
	names = append(names, personName{"ABC", "D", "222*3*"})
	names = append(names, personName{"ABE", "C", "223*2*"})
	names = append(names, personName{"ABE", "D", "223*3*"})
	names = append(names, personName{"ABE", "G", "223*4*"})
	names = append(names, personName{"ABE", "J", "223*5*"})
	names = append(names, personName{"ABI", "A", "224*2*"})
	names = append(names, personName{"ABI", "D", "224*3*"})
	names = append(names, personName{"ABI", "G", "224*4*"})
	names = append(names, personName{"ABI", "J", "224*5*"})
	names = append(names, personName{"GEORGE", "C", "436743*2*"})
	names = append(names, personName{"LEI", "A", "534*2*"})
	names = append(names, personName{"LESK", "M", "5375*6*"})
	names = append(names, personName{"JIM", "M", "546*6*"})
	names = append(names, personName{"LIO", "N", "546*6*"})
	names = append(names, personName{"KIM", "O", "546*6*"})
	names = append(names, personName{"JKL", "M", "555*6*"})
	names = append(names, personName{"JKM", "M", "556*6*"})
	names = append(names, personName{"JMN", "M", "566*6*"})
	names = append(names, personName{"MNO", "P", "666*7*"})
	names = append(names, personName{"MNP", "P", "667*7*"})
	names = append(names, personName{"MPP", "P", "677*7*"})
	names = append(names, personName{"TOM", "C", "866*2*"})
	result := searchName("546*6*", names, 0, len(names)-1)
	fmt.Printf("result = %s\n", result)
	if len(result) != 3 {
		t.Error("result error.")
	}
}
