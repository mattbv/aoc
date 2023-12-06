package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	textString := readTextFile("011/input.txt")
	inputData := parseInputString(textString)
	fmt.Println(inputData)

	errorMargin := 1
	for i := range inputData["Time"] {
		time := inputData["Time"][i]
		distance := inputData["Distance"][i]

		low, high, err := findInterval(float64(time), float64(distance)+0.001)
		if err != nil {
			fmt.Println(err)
			continue
		}
		errorMargin *= int(high - low + 1)
		fmt.Printf("Time: %d, Distance: %d, Low: %d, High: %d\n", time, distance, int(low), int(high))
	}

	fmt.Printf("Number of possibilities: %d\n", errorMargin)
}

func parseInputString(textString string) map[string][]int {
	var inputMap map[string][]int
	inputMap = make(map[string][]int)

	lineString := strings.Split(textString, "\n")

	for _, line := range lineString {
		lineData := strings.Split(line, ":")
		if len(lineData) < 2 {
			continue
		}
		key := strings.TrimSpace(lineData[0])

		inputMap[key] = convertListToInt(strings.Split(lineData[1], " "))
	}

	return inputMap
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

func findInterval(t float64, d float64) (float64, float64, error) {
	var a float64 = 1
	var b float64 = -t
	var c float64 = d

	// Calculate the discriminant
	disc := b*b - 4*a*c

	switch {
	case disc < 0:
		return 0, 0, errors.New("No real roots.")
	case disc == 0:
		root := -b / (2 * a)
		if root > 0 && root < t {
			return root, root, errors.New("Only one valid root.")
		} else {
			return 0, 0, errors.New("No valid root.")
		}
	default:
		var root1 = (-b + math.Sqrt(disc)) / (2 * a)
		var root2 = (-b - math.Sqrt(disc)) / (2 * a)

		// Check if the roots fall within the range (0, t)
		if root1 < 1 || root1 >= t {
			root1 = -math.MaxFloat64
		}
		if root2 < 1 || root2 >= t {
			root2 = math.MaxFloat64
		}
		if root1 == -math.MaxFloat64 && root2 == math.MaxFloat64 {
			return 0, 0, errors.New("No valid roots.")
		} else {
			// Return floor of larger root and ceil of smaller root
			return math.Ceil(math.Min(root1, root2)), math.Floor(math.Max(root1, root2)), nil
		}
	}
}
