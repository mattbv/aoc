package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func main() {
	text, lineLength, err := readFile("005/input.txt")
	original_text := text
	if err != nil {
		log.Fatal(err)
	}

	charactersExp := regexp.MustCompile(`[^\d\.]`)
	for i := 0; i < 3; i++ {
		resCharacterIndex := charactersExp.FindAllStringIndex(text, -1)
		for _, characterIndex := range resCharacterIndex {
			for j := -1; j <= 1; j++ {
				for k := -1; k <= 1; k++ {
					targetIndex := characterIndex[0] + (j * lineLength) + k
					if targetIndex < 0 || targetIndex >= len(text) {
						continue
					}
					if unicode.IsDigit(rune(text[targetIndex])) {
						text = replaceAtIndex(text, '*', targetIndex)
					}
				}
			}

		}
	}

	digitsExp := regexp.MustCompile(`\d+`)
	resOriginalDigits := digitsExp.FindAllString(original_text, -1)
	resDigits := digitsExp.FindAllString(text, -1)

	originalSum := getListSum(resOriginalDigits)
	sum := getListSum(resDigits)
	diff := originalSum - sum

	fmt.Println(diff)

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
