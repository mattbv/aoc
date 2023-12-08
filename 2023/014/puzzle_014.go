package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var reMap = strings.NewReplacer(
	"J", "A",
	"2", "B",
	"3", "C",
	"4", "D",
	"5", "E",
	"6", "F",
	"7", "G",
	"8", "H",
	"9", "I",
	"T", "J",
	"Q", "L",
	"K", "M",
	"A", "N",
)

func main() {
	file, err := os.Open("013/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cards := make(map[string]int)
	hand := make(map[string]int)
	for scanner.Scan() {
		lineString := scanner.Text()
		listFields := strings.Fields(lineString)

		newCards := reMap.Replace(listFields[0])

		counter := make(map[string]int)
		for _, c := range newCards {
			counter[string(c)]++
		}
		sortedCards := sortByValues(counter)

		counterList := make([]int, 0)
		aCounter := 0
		for _, key := range sortedCards {
			if key == "A" {
				aCounter += counter[string(key)]
				continue
			}
			value := counter[string(key)]
			counterList = append(counterList, value)
		}

		if len(counterList) == 0 {
			counterList = append(counterList, 5)
		} else {
			counterList[0] += aCounter
		}

		handScore := scoreHand(counterList)
		newKey := reMap.Replace(strconv.Itoa(handScore)) + newCards
		hand[newKey] = scoreHand(counterList)
		cards[newKey], err = strconv.Atoi(listFields[1])
	}

	finalSortedCards := sortByKeys(cards)
	finalScore := 0
	for rank, key := range finalSortedCards {
		finalScore += cards[key] * (rank + 1)
	}

	fmt.Printf("Final score: %d\n", finalScore)
}

func scoreHand(counterList []int) int {

	switch {
	case counterList[0] == 5:
		return 7
	case counterList[0] == 4:
		return 6
	case counterList[0] == 3 && counterList[1] == 2:
		return 5
	case counterList[0] == 3:
		return 4
	case counterList[0] == 2 && counterList[1] == 2:
		return 3
	case counterList[0] == 2:
		return 2
	default:
		return 1
	}
}

func sortByKeys(cards map[string]int) []string {
	keys := make([]string, 0, len(cards))

	for key := range cards {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	return keys
}

func sortByValues(cards map[string]int) []string {
	keys := make([]string, 0, len(cards))

	for key := range cards {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return cards[keys[i]] < cards[keys[j]]
	})

	ReverseSlice(keys)

	return keys
}

func ReverseSlice[T comparable](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return i > j
	})
}

func readTextFile(filename string) string {
	textBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	textString := string(textBytes)

	return textString
}
