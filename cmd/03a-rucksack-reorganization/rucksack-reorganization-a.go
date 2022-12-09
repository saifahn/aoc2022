package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// read line
	// check it is divisible by 2
	// split it into 2
	file, err := os.Open("data/03-test")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	var items []byte

	for sc.Scan() {
		line := sc.Text()
		if len(line)%2 != 0 {
			log.Fatal("error with input - the line has an uneven number of characters")
		}
		halfLength := len(line) / 2
		first := line[:halfLength]
		second := line[halfLength:]

		itemsMap := map[byte]int{}
		for i := 0; i < len(first); i++ {
			current := first[i]
			itemsMap[current] += 1
		}
		for j := 0; j < len(second); j++ {
			current := second[j]
			if itemsMap[current] != 0 {
				items = append(items, current)
				break
			}
		}
	}

	totalPriority := 0
	for i := 0; i < len(items); i++ {
		totalPriority += assignValueToLetter(items[i])
	}

	fmt.Println(items)
	fmt.Println(totalPriority)
}

func assignValueToLetter(letter byte) int {
	value := int(letter)
	// 97 - 122 is a-z
	value -= 96
	if value < 0 {
		// 65 - 90 is A-Z
		// 32 between 65 and 97, but we want to add 26 onto that value as it should
		// be afterwards
		value += 32 + 26
	}
	fmt.Println(value)
	return value
}
