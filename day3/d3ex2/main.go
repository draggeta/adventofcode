package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Number struct {
	Number int
	XMin   int
	XMax   int
	Y      int
}
type Gear struct {
	Gear      string
	Ratio     int
	Adjacency int
	X         int
	Y         int
}

// regex for numbers
var ren = regexp.MustCompile(`\d+`)

// regex for gears
var reg = regexp.MustCompile(`\*`)

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

func findGears(s string, y int, gears []Gear) []Gear {
	gearIndices := reg.FindAllStringIndex(s, -1)

	for _, i := range gearIndices {
		match := s[i[0]:i[1]]
		x := i[0]

		gear := Gear{
			Gear:  match,
			Ratio: 1,
			X:     x,
			Y:     y,
		}

		gears = append(gears, gear)
	}
	return gears
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
	gears := make([]Gear, 0)

	var y int = 0

	for scanner.Scan() {
		numbers = findNumbers(scanner.Text(), y, numbers)
		gears = findGears(scanner.Text(), y, gears)
		y++
	}

	// Not very efficient, but helps with debugging.
	var gearSlice []int
	for idx, i := range gears {
		for _, j := range numbers {
			if (i.Y-1 > j.Y) || (i.Y+1 < j.Y) {
				continue
			}
			if (i.X < j.XMin-1) || (i.X > j.XMax+1) {
				continue
			}
			gears[idx].Ratio *= j.Number
			gears[idx].Adjacency += 1
		}
		if gears[idx].Adjacency == 2 {
			gearSlice = append(gearSlice, gears[idx].Ratio)
		}
	}

	sum := 0
	for i := 0; i < len(gearSlice); i++ {
		sum += gearSlice[i]
	}

	fmt.Println(numbers)
	fmt.Println(gears)
	fmt.Println(gearSlice)
	fmt.Println(sum)
}
