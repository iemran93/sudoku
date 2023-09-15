package main

import (
	"os"
	"github.com/01-edu/z01"
)

var args []string = os.Args

func main() {

	solve()

	for i := range args[1:] {
		line := args[i]
		for _, chr := range line {
			z01.PrintRune(chr)
		}
	}
}

// possible
func possible(y, x, n int) bool {
	args = args[1:]
	// find if it in horizantal (x)
	// range of x index
	for i := 0; i < 9; i++{
		if args[y][i] == rune(n) {
			return false
		}
	}
	// find if it in vertical (y)
	// range of y index
	for i := 0; i < 9; i++{
		if args[i][x] == rune(n) {
			return false
		}
	}
	// find if it in square
	// first box of the square (nw) 
	xx := (x/3) * 3
	yy := (y/3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if args[yy+i][xx+j] == rune(n) {
				return false
			}
		}
	}
	return true
}

// solve
func solve() {
		// list of inputs(strings)
		args = args[1:]
		// range over y(columns)
		for i := range args {
			y := args[i]
			// range over x(rows)
			for ind, chr := range y {
				// if find . (empty)
				if chr == '.' {
					// try (i==y, ind==x, n==number)
					for n := '0'; i <= '9'; i++ {
						if possible(i, ind, int(n - 48)) {
							args[i][ind] = n
							solve()
							args[i][ind] = '.'
						}
					}
				}
			}
		}
}