package main

import (
	"fmt"
)

func strLen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func main() {

	fmt.Println(strLen("Hello world!"))
}
