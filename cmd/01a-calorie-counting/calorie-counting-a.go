package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("data/01-real")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	highest := 0
	currentTotal := 0

	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			if currentTotal > highest {
				highest = currentTotal
			}
			currentTotal = 0
			continue
		}
		item, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		currentTotal += item
	}

	println("The result is", highest)
}
