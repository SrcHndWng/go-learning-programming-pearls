package utils

import "sort"
import "strings"

// SortStr sorts alphabets in str.
func SortStr(str string) string {
	slice := make([]string, 0)
	for i := 0; i < len(str); i++ {
		slice = append(slice, str[i:i+1])
	}
	sort.Strings(slice)
	return strings.Join(slice, "")
}
