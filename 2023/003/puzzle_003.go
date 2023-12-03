package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

func main() {
	file, err := os.Open("003/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int
	for scanner.Scan() {
		text := scanner.Text()

		gameId, gameInfo, err := detectMaxCubes(text)
		if err != nil {
			log.Fatal(err)
		}

		if gameInfo["red"] <= maxRed && gameInfo["green"] <= maxGreen && gameInfo["blue"] <= maxBlue {
			sum += gameId
		}
		fmt.Printf("%s -> %d\n", text, sum)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}

func detectMaxCubes(text string) (int, map[string]int, error) {
	maxCubeColors := make(map[string]int)

	gameInfo := strings.Split(text, ": ")
	id, err := strconv.Atoi(strings.Split(gameInfo[0], "Game ")[1])
	if err != nil {
		return id, maxCubeColors, err
	}

	for _, subsetData := range strings.Split(gameInfo[1], "; ") {
		subsetInfo := strings.Split(subsetData, ", ")
		for _, cubeData := range subsetInfo {
			cubeInfo := strings.Split(cubeData, " ")
			cubeCount, err := strconv.Atoi(cubeInfo[0])
			if err != nil {
				continue
			}
			if cubeCount > maxCubeColors[cubeInfo[1]] {
				maxCubeColors[cubeInfo[1]] = cubeCount
			}
		}
	}

	return id, maxCubeColors, nil
}
