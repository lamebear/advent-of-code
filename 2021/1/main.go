package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/lamebear/advent-of-code/utils"
)

func main() {
	prior, increase := 0, 0
	depths := make([]int, 0, 2000)
	utils.ForInput("input.txt", func(line string, first bool) {
		current, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		depths = append(depths, current)
		if first {
			prior = current
			first = false
			return
		}

		if current > prior {
			increase += 1
		}

		prior = current
	})

	fmt.Printf("Increase Count: %d\n", increase)

	prior = 0
	increase = 0
	for i := 0; i < len(depths)-2; i++ {
		sum := depths[i] + depths[i+1] + depths[i+2]
		if i == 0 {
			prior = sum
			continue
		}

		if sum > prior {
			increase += 1
		}

		prior = sum
	}

	fmt.Printf("Sum Increase Count: %d\n", increase)
}
