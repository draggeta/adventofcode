package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const Red = 12
const Green = 13
const Blue = 14

type Game struct {
	Id    int
	Red   int
	Green int
	Blue  int
}

func (g Game) IsValid() bool {
	if g.Red > Red {
		fmt.Printf("Game %d color 'Red' with value %d exceeds max value %d\n", g.Id, g.Red, Red)
		return false
	}
	if g.Green > Green {
		fmt.Printf("Game %d color 'Green' with value %d exceeds max value %d\n", g.Id, g.Green, Green)
		return false
	}
	if g.Blue > Blue {
		fmt.Printf("Game %d color 'Blue' with value %d exceeds max value %d\n", g.Id, g.Blue, Blue)
		return false
	}
	return true
}

func gameParser(s string) int {
	idString, t, ok := strings.Cut(s, ":")
	if !ok {
		fmt.Println("No separator found!")
	}

	_, idString2, _ := strings.Cut(idString, " ")
	id, err := strconv.Atoi(idString2)
	if err != nil {
		fmt.Println(err)
	}

	for _, set := range strings.Split(t, ";") {
		game := Game{
			Id:    id,
			Red:   0,
			Green: 0,
			Blue:  0,
		}
		for _, g := range strings.Split(strings.TrimSpace(set), ",") {
			countStr, color, _ := strings.Cut(strings.TrimSpace(g), " ")
			countInt, err := strconv.Atoi(countStr)
			if err != nil {
				fmt.Println("Die count could not be converted.")
			}

			switch color {
			case "red":
				game.Red += countInt
			case "green":
				game.Green += countInt
			case "blue":
				game.Blue += countInt
			}
		}
		if !game.IsValid() {
			return 0
		}
	}
	return id
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int
	for scanner.Scan() {
		sum += gameParser(scanner.Text())
	}
	fmt.Println(sum)
}
