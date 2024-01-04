package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"sync"
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
	var wg sync.WaitGroup

	lines := make(chan string)
	coordinates := make(chan int)

	file, err := os.Open("./input.txt")
	if err != nil {
		return 0, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\d`)

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() error {
			defer wg.Done()

			for line := range lines {
				numericChars := re.FindAllString(line, -1)
				firstDigit := numericChars[0]
				lastDigit := numericChars[len(numericChars)-1]

				value := fmt.Sprintf("%s%s", firstDigit, lastDigit)

				val, err := strconv.Atoi(value)
				if err != nil {
					return err
				}

				coordinates <- val

			}

			return nil
		}()
	}

	go func() {
		defer close(lines)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
	}()

	go func() {
		wg.Wait()
		close(coordinates)
	}()

	var total int
	for c := range coordinates {
		total += c
	}

	return total, nil
}
