package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("data/01-real")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	elves := []int{}

	current := 0
	for sc.Scan() {
		line := sc.Text()
		// fmt.Println("line", line)
		if line == "" {
			elves = append(elves, current)
			current = 0
			continue
		}
		item, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		current += item
	}
	// account for the final line
	elves = append(elves, current)

	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})

	// fmt.Printf("%v\n", elves)

	fmt.Println(elves[0] + elves[1] + elves[2])
}
