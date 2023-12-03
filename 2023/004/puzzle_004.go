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
	file, err := os.Open("003/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int
	for scanner.Scan() {
		text := scanner.Text()

		_, gameInfo, err := detectMaxCubes(text)
		if err != nil {
			log.Fatal(err)
		}

		var powerCubes int = 1
		for _, cubeCount := range gameInfo {
			powerCubes *= cubeCount
		}

		sum += powerCubes

		fmt.Printf("%s -> %d\n", text, powerCubes)

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
