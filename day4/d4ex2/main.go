package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Card struct {
	Card           string
	copies         int
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

	return score
}

func parseCard(s string) Card {
	cardTxt, numbersTxt, _ := strings.Cut(s, ":")
	_, cardStr, _ := strings.Cut(cardTxt, " ")

	winTxt, numTxt, _ := strings.Cut(numbersTxt, "|")

	winSlice := strings.Fields(strings.TrimSpace(winTxt))
	numSlice := strings.Fields(strings.TrimSpace(numTxt))

	return Card{Card: cardStr, copies: 1, winningNumbers: winSlice, numbers: numSlice}
}

func main() {
	filePath := "../input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("File '%s' not found", filePath)
		os.Exit(1)
	}
	defer file.Close()

	cardList := make([]Card, 198)

	scanner := bufio.NewScanner(file)
	idx := 0
	for scanner.Scan() {
		cardList[idx] = parseCard(scanner.Text())
		idx++
	}

	// instead of inserting a copy each time, increment the amount of copies you have.
	// I don't know if it is faster, but the list is significantly smaller.
	sum := 0
	for i := 0; i < len(cardList); i++ {
		card := cardList[i]
		sum += card.copies
		score := card.Score()
		for j := i + 1; j < i+score+1 && j < len(cardList); j++ {
			cardList[j].copies += card.copies
		}
		fmt.Println(card.Card, card.copies, score)
	}

	fmt.Printf("Cards total: %d\n", sum)
}
