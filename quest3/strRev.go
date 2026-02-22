package main

import (
	"fmt"
)

func strRev(s string) string {
	runes := []rune(s)
	result := make([]rune, 0, len(runes))

	for i := len(runes) - 1; i >= 0; i-- {
		result = append(result, runes[i])
	}
	return string(result)
}

func main() {
	fmt.Println(strRev("Hello world!"))
}
