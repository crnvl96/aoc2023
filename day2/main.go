package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func main() {
	res, err := cubeConundrum()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}

func cubeConundrum() (int, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return 0, err
	}

	defer file.Close()

	type threshold struct {
		re  *regexp.Regexp
		val int
	}

	thresholds := map[string]threshold{
		"green": {
			re:  regexp.MustCompile(`\d+\ green`),
			val: 13,
		},
		"red": {
			re:  regexp.MustCompile(`\d+\ red`),
			val: 12,
		},
		"blue": {
			re:  regexp.MustCompile(`\d+\ blue`),
			val: 14,
		},
	}

	type data struct {
		index int
		line  string
	}

	dataCh := make(chan data)

	var wg sync.WaitGroup

	scanner := bufio.NewScanner(file)

	go func() {
		defer close(dataCh)

		index := 0
		for scanner.Scan() {
			index++
			data := data{
				index: index,
				line:  scanner.Text(),
			}

			dataCh <- data
		}
	}()

	totalCh := make(chan int)

	threadsNo := 10

	for i := 0; i < threadsNo; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

		scanner:
			for data := range dataCh {
				sample := strings.Split(data.line, ":")

				for k := range thresholds {
					matches := thresholds[k].re.FindAllString(sample[1], -1)

					for i := 0; i < len(matches); i++ {
						if val, err := strconv.Atoi(strings.Trim(matches[i], fmt.Sprintf(" %s ", k))); err == nil &&
							val > thresholds[k].val {
							continue scanner
						}
					}
				}

				totalCh <- data.index

			}
		}()
	}

	go func() {
		wg.Wait()
		close(totalCh)
	}()

	var total int

	for t := range totalCh {
		total += t
	}

	return total, nil
}
