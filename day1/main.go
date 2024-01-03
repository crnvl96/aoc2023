package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	total, err := trebuchet()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(total)
}

func trebuchet() (int, error) {
	var coordinates []int
	var total int

	file, err := os.Open("./input.txt")
	if err != nil {
		return 0, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\d`)

	for scanner.Scan() {
		numericChars := re.FindAllString(scanner.Text(), -1)
		firstDigit := numericChars[0]
		lastDigit := numericChars[len(numericChars)-1]

		coordinate := fmt.Sprintf("%s%s", firstDigit, lastDigit)

		val, err := strconv.Atoi(coordinate)
		if err != nil {
			return 0, err
		}

		coordinates = append(coordinates, val)
	}

	for _, c := range coordinates {
		total += c
	}

	return total, nil
}
