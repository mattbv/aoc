package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
		re := regexp.MustCompile(`\d`)
		text := scanner.Text()
		matches := re.FindAllString(text, -1)

		first, last := matches[0], matches[len(matches)-1]

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
