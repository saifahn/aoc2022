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
	file, err := os.Open("data/02-real")
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
		outcome := split[1]
		points += outcomePoint(outcome)
		shape := calculateShapeBasedOnVillainAndOutcome(villain, outcome)
		points += shapePointMap[shape]
	}

	fmt.Println(points)
}

var outcomeLetterStringMap = map[string]string{
	"X": "Lose",
	"Y": "Draw",
	"Z": "Win",
}

var outcomePointMap = map[string]int{
	"Lose": 0,
	"Draw": 3,
	"Win":  6,
}

func outcomePoint(outcomeLetter string) int {
	outcome := outcomeLetterStringMap[outcomeLetter]
	return outcomePointMap[outcome]
}

var villainLetterStringMap = map[string]string{
	"A": "Rock",
	"B": "Paper",
	"C": "Scissors",
}

var villainOutcomeShapeMap = map[string]map[string]string{
	"Rock": {
		"Lose": "Scissors",
		"Draw": "Rock",
		"Win":  "Paper",
	},
	"Paper": {
		"Lose": "Rock",
		"Draw": "Paper",
		"Win":  "Scissors",
	},
	"Scissors": {
		"Lose": "Paper",
		"Draw": "Scissors",
		"Win":  "Rock",
	},
}

func calculateShapeBasedOnVillainAndOutcome(villainLetter, outcomeLetter string) string {
	villain := villainLetterStringMap[villainLetter]
	outcome := outcomeLetterStringMap[outcomeLetter]
	return villainOutcomeShapeMap[villain][outcome]
}

var shapePointMap = map[string]int{
	"Rock":     1,
	"Paper":    2,
	"Scissors": 3,
}
