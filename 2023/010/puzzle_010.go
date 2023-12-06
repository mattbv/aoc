package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	textString := readTextFile("009/input.txt")
	sectionSequence := [7]string{
		"seed-to-soil map",
		"soil-to-fertilizer map",
		"fertilizer-to-water map",
		"water-to-light map",
		"light-to-temperature map",
		"temperature-to-humidity map",
		"humidity-to-location map",
	}

	parsedData := structuredDataFromString(textString)
	visitedStates := make(map[int]map[int]bool)

	locationIds := make([]int, 0)

	seedsSet := make(map[int]struct{})
	for i := 0; i < len(parsedData["seeds"][0]); i += 2 {
		for j := 0; j < parsedData["seeds"][0][i+1]; j++ {
			seedsSet[parsedData["seeds"][0][i]+j] = struct{}{}
		}
	}
	fmt.Printf("Starting with %d seeds\n", len(seedsSet))

	// TODO: This is a bit of a mess, but it works. The better option would be to check entire ranges against each other.
	for seed, _ := range seedsSet {
		fmt.Printf("Processing with seed %d\n", seed)
		startId := seed
		for stepId, sequenceStep := range sectionSequence {
		seqLoop:
			for _, stepData := range parsedData[sequenceStep] {
				if !visitedStates[stepId][startId] {
					diff := startId - stepData[1]
					if diff >= 0 && diff < stepData[2] {
						startId = stepData[0] + diff
						if visitedStates[stepId] == nil {
							visitedStates[stepId] = make(map[int]bool)
						}
						visitedStates[stepId][startId] = true
						break seqLoop
					}
				}
			}
		}
		locationIds = append(locationIds, startId)
	}
	fmt.Println(locationIds)
	fmt.Printf("The minimum location id is %d\n", min(locationIds))
}

func min(list []int) int {
	minValue := list[0]
	for _, item := range list {
		if item < minValue {
			minValue = item
		}
	}
	return minValue
}

func structuredDataFromString(textString string) map[string][][]int {
	parsedData := make(map[string][][]int)

	textSectionsString := strings.Split(textString, "\n\n")
	for _, sectionString := range textSectionsString {
		sectionDataString := strings.Split(sectionString, ":")
		sectionTitle := sectionDataString[0]
		sectionMapString := strings.Split(sectionDataString[1], "\n")

		sectionData := make([][]int, 0)
		for _, sectionRow := range sectionMapString {
			if sectionRow == "" {
				continue
			}
			rowValuesString := strings.Split(sectionRow, " ")
			rowValues := convertListToInt(rowValuesString)
			sectionData = append(sectionData, rowValues)
		}

		parsedData[sectionTitle] = sectionData
	}

	return parsedData
}

func readTextFile(filename string) string {
	textBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	textString := string(textBytes)

	return textString
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
