package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// scan each line
	// each line, scan the chars into vars
	file, err := os.Open("data/02-test")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	points := 0

	// scan over the lines
	for sc.Scan() {
		line := sc.Text()
		split := strings.Split(line, " ")
		villain := split[0]
		hero := split[1]
		points += shapePoints(hero)
		points += outcomePoints(villain, hero)
	}

	fmt.Println(points)
}

var SHAPE_POINT = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func shapePoints(shape string) int {
	return SHAPE_POINT[shape]
}

var OUTCOME_POINTS = map[string]map[string]int{
	"A": {
		"X": 3,
		"Y": 6,
		"Z": 0,
	},
	"B": {
		"X": 0,
		"Y": 3,
		"Z": 6,
	},
	"C": {
		"X": 6,
		"Y": 0,
		"Z": 3,
	},
}

// takes the two shapes (villain must be A, B, C)
// hero must be X, Y, Z
// calculates the outcome score
func outcomePoints(villain, hero string) int {
	return OUTCOME_POINTS[villain][hero]
}
