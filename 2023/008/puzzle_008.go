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
	file, err := os.Open("007/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cardRegistery := make(map[int]int)
	for scanner.Scan() {
		test := scanner.Text()
		cardId, winningNumbers, ownNumbers, err := getGameInfo(test)

		if err != nil {
			log.Fatal(err)
			continue
		}

		matchCount := processCard(winningNumbers, ownNumbers)
		cardRegistery[cardId] += 1

		for j := 1; j <= matchCount; j++ {
			cardRegistery[j+cardId] += cardRegistery[cardId]
		}

		fmt.Printf("Card %d: %d\n", cardId, matchCount)
	}

	fmt.Println(cardRegistery)

	sum := 0
	for _, value := range cardRegistery {
		sum += value
	}

	fmt.Printf("Sum: %d\n", sum)
}

func processCard(winningNumbers []int, ownNumbers []int) int {
	matchCount := 0
	for _, ownNumber := range ownNumbers {
		for _, winningNumber := range winningNumbers {
			if ownNumber == winningNumber {
				matchCount++
			}
		}
	}
	return matchCount
}

func getGameInfo(text string) (int, []int, []int, error) {
	var winningNumbers []int
	var ownNumbers []int

	cardInfo := strings.Split(text, ":")
	cardIdString := strings.TrimSpace(strings.Split(cardInfo[0], "Card ")[1])
	cardId, err := strconv.Atoi(cardIdString)

	numbersString := strings.Split(cardInfo[1], "|")
	winningNumbersString := strings.Split(numbersString[0], " ")
	ownNumbersString := strings.Split(numbersString[1], " ")

	winningNumbers = convertListToInt(winningNumbersString)
	ownNumbers = convertListToInt(ownNumbersString)

	return cardId, winningNumbers, ownNumbers, err
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
