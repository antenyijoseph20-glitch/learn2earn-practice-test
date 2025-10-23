package main

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, word := range a {
		for _, r := range word {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

func main() {
	PrintWordsTables([]string{"Hello", "World", "Piscine"})
}
