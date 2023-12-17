package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Number struct {
	Number    int
	Adjacency int
	XMin      int
	XMax      int
	Y         int
}
type Symbol struct {
	Symbol string
	X      int
	Y      int
}

// regex for numbers
var ren = regexp.MustCompile(`\d+`)

// regex for symbols
var res = regexp.MustCompile(`[^\.\d\s]`)

func findNumbers(s string, y int, numbers []Number) []Number {
	numIndices := ren.FindAllStringIndex(s, -1)

	for _, i := range numIndices {
		match, err := strconv.Atoi(s[i[0]:i[1]])
		if err != nil {
			fmt.Println("Match isn't an integer.")
		}
		xmin, xmax := i[0], i[1]-1

		num := Number{
			Number: match,
			XMin:   xmin,
			XMax:   xmax,
			Y:      y,
		}

		numbers = append(numbers, num)
	}
	return numbers
}

func findSymbols(s string, y int, symbols []Symbol) []Symbol {
	symIndices := res.FindAllStringIndex(s, -1)

	for _, i := range symIndices {
		match := s[i[0]:i[1]]
		x := i[0]

		sym := Symbol{
			Symbol: match,
			X:      x,
			Y:      y,
		}

		symbols = append(symbols, sym)
	}
	return symbols
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// var str string = "lol"

	numbers := make([]Number, 0)
	symbols := make([]Symbol, 0)

	var y int = 0

	for scanner.Scan() {
		numbers = findNumbers(scanner.Text(), y, numbers)
		symbols = findSymbols(scanner.Text(), y, symbols)
		y++
	}

	// Not very efficient, but helps with debugging.
	var numSlice []int
	for _, i := range symbols {
		for idx, j := range numbers {
			if (i.Y-1 > j.Y) || (i.Y+1 < j.Y) {
				continue
			}
			if (i.X < j.XMin-1) || (i.X > j.XMax+1) {
				continue
			}
			if j.Adjacency != 0 {
				continue
			}
			numbers[idx].Adjacency += 1
			numSlice = append(numSlice, j.Number)
		}
	}

	sum := 0
	for i := 0; i < len(numSlice); i++ {
		sum += numSlice[i]
	}

	fmt.Println(numbers)
	fmt.Println(symbols)
	fmt.Println(numSlice)
	fmt.Println(sum)
}
