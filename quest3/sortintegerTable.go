package main

import (
	"fmt"
)

func main() {
	s := []int{5, 4, 3, 2, 1, 0}

	fmt.Println(SortIntegerTable(s))
}

func SortIntegerTable(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}
