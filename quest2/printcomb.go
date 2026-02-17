package main

import "github.com/01-edu/z01"

func main() {
	printComb()
}

func printComb() {
	firstComb := true //end of first comb
	for i := '0'; i < '9'; i++ {
		for j := '0'; j < '9'; j++ {
			for k := '0'; k < '9'; k++ {
				if k > j && j > i {
					if !firstComb {
						z01.PrintRune(',')
						z01.PrintRune(' ')

					}
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					firstComb = false

				}
			}
		}
	}
	z01.PrintRune('\n')
}
