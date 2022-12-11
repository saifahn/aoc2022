package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data/03-test")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	var items []byte

	elfOfGroup := 1
	// 0 is not yet, 1 is exists in first
	// 2 is exists in second
	itemMap := map[byte]int{}

	for sc.Scan() {
		line := sc.Text()

		for i := 0; i < len(line); i++ {
			current := line[i]
			// mark the item as in the first line
			if elfOfGroup == 1 && itemMap[current] == 0 {
				itemMap[current] = 1
			}
			// mark the item as in the second line as well as first line
			if elfOfGroup == 2 && itemMap[current] == 1 {
				itemMap[current] = 2
			}
			// if in the third line, add to items, reset the map and continue out of here
			if elfOfGroup == 3 && itemMap[current] == 2 {
				items = append(items, current)
				itemMap = map[byte]int{}
				continue
			}
			if i == len(line)-1 {
				// go to the next elf of group
				elfOfGroup += 1
				if elfOfGroup > 3 {
					elfOfGroup = 1
					itemMap = map[byte]int{}
				}
			}
		}
	}
	fmt.Printf("%s", items)
}
