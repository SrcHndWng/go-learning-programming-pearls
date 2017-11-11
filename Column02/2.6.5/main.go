package main

import "flag"
import "fmt"

func main() {
	flag.Parse()
	str := flag.Args()[0]
	input := makeSlice(str)
	fmt.Printf("input = %s\n", input)
	result := reverse(input)
	fmt.Printf("result = %v\n", result)
}

func reverse(slice []string) []string {
	result := make([]string, 0)
	for i := 0; i < len(slice); i++ {
		result = append(result, slice[len(slice)-1-i])
	}
	return result
}

func makeSlice(str string) []string {
	result := make([]string, 0)
	for i := 0; i < len(str); i++ {
		result = append(result, str[i:i+1])
	}
	return result
}
