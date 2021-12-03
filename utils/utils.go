package utils

import (
	"bufio"
	"log"
	"os"
)

func ForInput(path string, processor func(line string, first bool)) {
	file, err := os.Open(path)
	HandleErr(err)

	first := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		processor(scanner.Text(), first)
		first = false
	}
}

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
