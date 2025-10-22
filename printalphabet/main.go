package main

import "github.com/01-edu/z01"

func main() {
	for r := 'A'; r <= 'Z'; r++ {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
