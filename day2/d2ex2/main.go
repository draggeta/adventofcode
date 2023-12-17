package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MinGameDie struct {
	Id    int
	Red   int
	Green int
	Blue  int
}

func gameParser(s string) MinGameDie {
	idString, t, ok := strings.Cut(s, ":")
	if !ok {
		fmt.Println("No separator found!")
	}

	_, idString2, _ := strings.Cut(idString, " ")
	id, err := strconv.Atoi(idString2)
	if err != nil {
		fmt.Println(err)
	}

	game := MinGameDie{
		Id: id,
		// Red:   0,
		// Green: 0,
		// Blue:  0,
	}
	for _, set := range strings.Split(t, ";") {
		for _, cube := range strings.Split(strings.TrimSpace(set), ",") {
			number, color := getCube(cube)

			switch color {
			case "red":
				if game.Red < number {
					game.Red = number
				}
			case "green":
				if game.Green < number {
					game.Green = number
				}
			case "blue":
				if game.Blue < number {
					game.Blue = number
				}
			}
		}
	}

	return game
}

func getCube(cube string) (int, string) {
	cube = strings.TrimSpace(cube)
	numberStr, color, _ := strings.Cut(cube, " ")
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		fmt.Println("Die count could not be converted.")
	}

	return number, color
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sumPower int
	for scanner.Scan() {
		game := gameParser(scanner.Text())
		r, g, b := game.Red, game.Green, game.Blue
		setPower := r * g * b
		sumPower += setPower
	}
	fmt.Println(sumPower)
}
