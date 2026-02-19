package main

import (
	"github.com/01-edu/z01"
)

func main() {
	printNbr(-123)
	printNbr(0)
	printNbr(123)
	z01.PrintRune('\n')
}

func printNbr(n int) {
	if n == -9223372036854775808 {
		for _, r := range "-9223372036854775808" {
			z01.PrintRune(r)
		}
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	if n == 0 {
		z01.PrintRune('0')
	}

	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(rune('0' + digits[i]))
	}
}
