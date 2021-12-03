package main

import (
	"github.com/lamebear/advent-of-code/utils"
)

func main() {
	values := make([]string, 0, 2000)
	utils.ForInput("input.txt", func(line string, first bool) {
		values = append(values, line)
	})

	for _, value := range values {

	}
}
