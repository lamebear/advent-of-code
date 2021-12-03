package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lamebear/advent-of-code/utils"
)

func main() {
	values := make([]string, 0, 2000)
	forward, depth, aim := 0, 0, 0
	utils.ForInput("input.txt", func(line string, first bool) {
		values = append(values, line)
	})

	for _, value := range values {
		commands := strings.Split(value, " ")
		num, err := strconv.Atoi(commands[1])
		utils.HandleErr(err)

		switch commands[0] {
		case "up":
			aim -= num
		case "down":
			aim += num
		case "forward":
			forward += num
			depth += aim * num
		}

		if depth < 0 {
			depth = 0
		}
	}

	fmt.Println(forward * depth)
}
