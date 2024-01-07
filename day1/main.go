package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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
	path, err := filepath.Abs("./input.txt")
	if err != nil {
		return 0, err
	}

	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	re := regexp.MustCompile(`\d`)

	lineCh := make(chan string)
	coordCh := make(chan int)

	var wg sync.WaitGroup

	threadsNo := 20

	scanner := bufio.NewScanner(file)

	go func() {
		defer close(lineCh)
		for scanner.Scan() {
			lineCh <- scanner.Text()
		}
	}()

	for i := 0; i < threadsNo; i++ {
		wg.Add(1)
		go func() error {
			defer wg.Done()

			for line := range lineCh {
				numericChars := re.FindAllString(line, -1)
				firstDigit := numericChars[0]
				lastDigit := numericChars[len(numericChars)-1]

				value := fmt.Sprintf("%s%s", firstDigit, lastDigit)

				val, err := strconv.Atoi(value)
				if err != nil {
					return err
				}

				coordCh <- val

			}

			return nil
		}()
	}

	go func() {
		wg.Wait()
		close(coordCh)
	}()

	var total int
	for c := range coordCh {
		total += c
	}

	return total, nil
}
