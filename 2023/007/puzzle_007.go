package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("007/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := float64(0)
	for scanner.Scan() {
		test := scanner.Text()
		matchCount := 0
		matchSum := float64(0)

		cardId, winningNumbers, ownNumbers := getGameInfo(test)
		fmt.Printf("Card ID: %s\n", cardId)
		matchCount, matchSum = processCard(winningNumbers, ownNumbers)
		fmt.Printf("Match sum: %f\n", matchSum)
		fmt.Printf("Match count: %d\n", matchCount)

		sum += matchSum

	}

	fmt.Printf("Sum: %f\n", sum)

}

func processCard(winningNumbers []int, ownNumbers []int) (int, float64) {
	matchCount := 0
	matchSum := float64(0)
	for _, ownNumber := range ownNumbers {
		for _, winningNumber := range winningNumbers {
			if ownNumber == winningNumber {
				matchCount++
			}
		}
	}
	fmt.Printf("Match count: %d\n", matchCount)
	if matchCount > 1 {
		matchSum += math.Pow(2, float64(matchCount-1))
	} else {
		matchSum += float64(matchCount)
	}

	return matchCount, matchSum
}

func getGameInfo(text string) (string, []int, []int) {
	var cardId string
	var winningNumbers []int
	var ownNumbers []int

	cardInfo := strings.Split(text, ":")
	cardId = strings.Split(cardInfo[0], " ")[1]

	numbersString := strings.Split(cardInfo[1], "|")
	winningNumbersString := strings.Split(numbersString[0], " ")
	ownNumbersString := strings.Split(numbersString[1], " ")

	winningNumbers = convertListToInt(winningNumbersString)
	ownNumbers = convertListToInt(ownNumbersString)

	return cardId, winningNumbers, ownNumbers
}

func convertListToInt(list []string) []int {
	var intList []int
	for _, item := range list {
		itemInt, err := strconv.Atoi(item)
		if err != nil {
			continue
		}
		intList = append(intList, itemInt)
	}
	return intList
}
