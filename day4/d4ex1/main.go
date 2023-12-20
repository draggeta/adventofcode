package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Card struct {
	Card           string
	winningNumbers []string
	numbers        []string
}

func contains(s []string, n string) bool {
	for _, i := range s {
		if i == n {
			return true
		}
	}
	return false
}

func (c *Card) Score() int {
	score := 0
	for _, i := range c.winningNumbers {
		if contains(c.numbers, i) {
			score += 1
		}
	}

	if score == 0 {
		return 0
	}
	return int(math.Pow(2, float64(score-1)))
}

func parseCard(s string) Card {
	cardTxt, numbersTxt, _ := strings.Cut(s, ":")
	_, cardStr, _ := strings.Cut(cardTxt, " ")

	winTxt, numTxt, _ := strings.Cut(numbersTxt, "|")

	winSlice := strings.Fields(strings.TrimSpace(winTxt))
	numSlice := strings.Fields(strings.TrimSpace(numTxt))

	return Card{Card: cardStr, winningNumbers: winSlice, numbers: numSlice}
}

func main() {
	filePath := "../input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("File '%s' not found", filePath)
		os.Exit(1)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		card := parseCard(scanner.Text())
		score := card.Score()
		fmt.Println(card.Card, score)
		sum += score
	}

	fmt.Printf("Points total: %d\n", sum)
}
