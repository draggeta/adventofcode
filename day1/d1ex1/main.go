package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func calibrationValue(s string) int {
	var first string
	var last string

	for _, c := range s {
		if !unicode.IsDigit(c) {
			continue
		}
		if first == "" {
			sc := string(c)
			first, last = sc, sc
			continue
		}
		last = string(c)

	}

	strCalVal := fmt.Sprintf("%s%s", first, last)

	calVal, err := strconv.Atoi(strCalVal)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(calVal)

	return calVal
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
		sum += calibrationValue(scanner.Text())
	}

	fmt.Printf("Calibration document value: %d\n", sum)
}
