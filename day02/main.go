package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gansach/AOC23/pkg/fileutil"
)

func main() {
	lines, _ := fileutil.ReadFileLines("inputs/day02.txt")
	A(lines)
	B(lines)
}

func A(games []string) {
	validGamesSum := 0

	maxCubesColorCntMap := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, game := range games {
		includeGame := true
		gameNum, turnsArr := processGame(game)
		maxGameColorCntMatch := generateMaxGameColorCntMap(turnsArr)

		for color, cnt := range maxGameColorCntMatch {
			if cnt > maxCubesColorCntMap[color] {
				includeGame = false
			}
			if !includeGame {
				break
			}
		}

		if includeGame {
			validGamesSum += gameNum
		}
	}
	fmt.Println(validGamesSum)
}

func B(games []string) {
	powerSum := 0
	for _, game := range games {
		gamePower := 1
		_, turns := processGame(game)
		maxGameColorCntMatch := generateMaxGameColorCntMap(turns)

		for _, cnt := range maxGameColorCntMatch {
			gamePower *= cnt
		}
		powerSum += gamePower
	}
	fmt.Println(powerSum)
}

func generateMaxGameColorCntMap(turns []string) map[string]int {
	maxGameColorCntMatch := make(map[string]int)
	for _, turn := range turns {
		cubesArr := strings.Split(turn, ",")

		for _, cube := range cubesArr {
			trimmedCube := strings.Trim(cube, " ")

			cubeRe := regexp.MustCompile(`(\d+) (.*)`)
			cubeMatches := cubeRe.FindStringSubmatch(trimmedCube)

			cnt, _ := strconv.Atoi(cubeMatches[1])
			color := cubeMatches[2]

			_, colorExists := maxGameColorCntMatch[color]
			if colorExists {
				maxGameColorCntMatch[color] = max(maxGameColorCntMatch[color], cnt)
			} else {
				maxGameColorCntMatch[color] = cnt
			}
		}
	}
	return maxGameColorCntMatch
}

func processGame(game string) (int, []string) {
	gameRe := regexp.MustCompile(`Game (\d+): (.*)`)
	gameMatches := gameRe.FindStringSubmatch(game)

	gameNum, _ := strconv.Atoi(gameMatches[1])
	turnsArr := strings.Split(gameMatches[2], ";")

	return gameNum, turnsArr
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
