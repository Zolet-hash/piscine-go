package main

import (
	"fmt"
)

func main() {
	a := 13
	b := 2
	ultimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)

}

func ultimateDivMod(a *int, b *int) {
	x := *a
	*a = *a / *b
	*b = x % *b
}
