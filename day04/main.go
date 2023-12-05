package main

import (
	"fmt"
	"regexp"

	"github.com/gansach/AOC23/pkg/defaultvaluemap"
	"github.com/gansach/AOC23/pkg/fileutil"
	"github.com/gansach/AOC23/pkg/math"
)

func main() {
	lines, _ := fileutil.ReadFileLines("inputs/day04.txt")
	A(lines)
	B(lines)
}

func A(cards []string) {
	sum := 0
	for _, card := range cards {
		cntOfNumsThatWon := getWinningNumsCount(card)
		if cntOfNumsThatWon > 0 {
			sum += math.Pow(2, cntOfNumsThatWon-1)
		}
	}
	fmt.Println(sum)
}

func B(cards []string) {
	cardsFreq := defaultvaluemap.New[int, int](0)

	for i, card := range cards {
		cardsFreq.Set(i, cardsFreq.Get(i)+1)

		cntOfNumsThatWon := getWinningNumsCount(card)
		for j := i + 1; j <= i+cntOfNumsThatWon; j++ {
			cardsFreq.Set(j, cardsFreq.Get(j)+cardsFreq.Get(i))
		}
	}

	totalFreq := 0
	for cardFreq := range cardsFreq.Range() {
		totalFreq += cardFreq.Value
	}
	fmt.Println(totalFreq)
}

func getWinningNumsCount(card string) int {
	cardRe := regexp.MustCompile(`Card\s+\d+:(.*)\|(.*)`)
	cardMatches := cardRe.FindStringSubmatch(card)

	winningNums := getAllNumsFromStr(cardMatches[1])
	winningNumsMap := map[string]bool{}
	for _, winningNum := range winningNums {
		winningNumsMap[winningNum] = true
	}

	nums := getAllNumsFromStr(cardMatches[2])
	cntOfNumsThatWon := 0
	for _, num := range nums {
		_, exists := winningNumsMap[num]
		if exists {
			cntOfNumsThatWon += 1
		}
	}
	return cntOfNumsThatWon
}

func getAllNumsFromStr(str string) []string {
	numRe := regexp.MustCompile(`\d+`)
	return numRe.FindAllString(str, -1)
}
