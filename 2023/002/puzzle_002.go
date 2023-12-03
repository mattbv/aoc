package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("001/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int
	for scanner.Scan() {
		text := scanner.Text()

		min_index := len(text)
		max_index := 0
		first := ""
		last := ""
		for _, number := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"} {
			first_index := strings.Index(text, number)
			last_index := strings.LastIndex(text, number)
			if first_index != -1 {
				if first_index <= min_index {
					min_index = first_index
					first = number
				}
			}
			if last_index != -1 {
				if last_index >= max_index {
					max_index = last_index
					last = number
				}
			}
		}

		first = convertToDigit(first)
		last = convertToDigit(last)
		doubleDigitStr := first + last
		doubleDigit, _ := strconv.Atoi(doubleDigitStr)
		fmt.Printf("%s -> %d\n", text, doubleDigit)
		sum += doubleDigit

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}

func convertToDigit(word string) string {
	switch word {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	}
	return word
}
