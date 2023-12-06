package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func firstCalibrationValue(s string) string {
	var first string
	for _, c := range s {
		if !unicode.IsDigit(c) {
			continue
		}
		first = string(c)
		break
	}

	return first
}

func reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func calibrationValue(s string) int {
	re := strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")

	v := re.Replace(s)
	first := firstCalibrationValue(v)

	rer := strings.NewReplacer("eno", "1", "owt", "2", "eerht", "3", "ruof", "4", "evif", "5", "xis", "6", "neves", "7", "thgie", "8", "enin", "9")
	sr := reverse(s)

	vr := rer.Replace(sr)
	last := firstCalibrationValue(vr)

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
