package main

import "fmt"

func UltimatePointer(n ***int) {
	***n = 1
}

func main() {
	a := 0
	b := &a
	n := &b
	UltimatePointer(&n)
	fmt.Println(a)
}