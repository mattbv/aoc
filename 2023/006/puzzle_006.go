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
	// text, lineLength, err := readFile("005/input.txt")
	text, lineLength, err := readFile("005/test.txt")
	fmt.Printf("Line length: %d\n", lineLength)
	if err != nil {
		log.Fatal(err)
	}

	gearExp := regexp.MustCompile(`\*`)
	resGearIndex := gearExp.FindAllStringIndex(text, -1)

	digitsExp := regexp.MustCompile(`\d+`)
	resDigitsIndex := digitsExp.FindAllStringIndex(text, -1)

	var sum int

	var overallNeighbours [][]string
	for _, gearIndex := range resGearIndex {
		var neighbours []string
		for _, digitsIndex := range resDigitsIndex {
			if itemInList(gearIndex[0], digitsIndex, lineLength) {
				neighbours = append(neighbours, text[digitsIndex[0]:digitsIndex[1]])
			}
			if len(neighbours) == 2 {
				overallNeighbours = append(overallNeighbours, neighbours)
				gearItem1, _ := strconv.Atoi(neighbours[0])
				gearItem2, _ := strconv.Atoi(neighbours[1])
				sum += gearItem1 * gearItem2
				break
			}
		}
	}
	fmt.Println(overallNeighbours)
	fmt.Println(sum)

}

func itemInList(item int, rangeList []int, lineLength int) bool {
	var check bool
	for i := 0; i < 3; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				targetIndex := item + (j * lineLength) + k
				if targetIndex >= rangeList[0] && targetIndex < rangeList[1] {
					check = true
				}
			}
		}
	}
	return check
}

func getListSum(list []string) int {
	var sum int
	for _, digit := range list {
		digitInt, err := strconv.Atoi(digit)
		if err != nil {
			continue
		}
		sum += digitInt
	}
	return sum
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func readFile(path string) (string, int, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return "", 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var text string
	var lineLength int
	for scanner.Scan() {
		currentLine := scanner.Text()
		lineLength = len(currentLine)
		text += currentLine
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return "", 0, err
	}

	return text, lineLength, nil
}
