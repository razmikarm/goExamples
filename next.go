package main

import (
	"fmt"
)

func next() {

	i := 0
	size := 9

	for i <= size {

		for j := 0; j <= size; j++ {
			cell := "* "
			if j != 0 && j != 9 && i != 0 && i != 9 && (j == 1 || j == 8 || i == 1 || i == 8) {
				// if i%2 == 0 && i == j {
				cell = "  "
			}
			fmt.Print(cell)
		}
		i++
		fmt.Println()
	}

}
