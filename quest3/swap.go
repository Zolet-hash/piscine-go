package main

import (
	"fmt"
)

func main() {
	a := 0
	b := 1
	swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}

func swap(a *int, b *int) {
	x := *a
	*a = *b
	*b = x
}
