package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("inputs/day06.txt")
	contentChunks := strings.Split(strings.TrimSpace(string(content)), "\r\n")

	raceTimes := getAllNumsFromStr(contentChunks[0])
	distanceTravelled := getAllNumsFromStr(contentChunks[1])

	A(raceTimes, distanceTravelled)

	B := getValidNumsCount(extractAndCombineDigits(contentChunks[0]), extractAndCombineDigits(contentChunks[1]))
	fmt.Println(B)
}

func A(raceTimes, distanceTravelled []int) {
	p := 1
	for i := range raceTimes {
		p *= getValidNumsCount(raceTimes[i], distanceTravelled[i])
	}
	fmt.Println(p)
}

func getValidNumsCount(raceTime, distanceToTravel int) int {
	t := float64(raceTime)
	d := float64(distanceToTravel)
	root1, root2 := getQuadraticRoots(1, -t, d)

	L := math.Floor(root1)
	S := math.Ceil(root2)

	validNums := L - S + 1

	if isInteger(root1) {
		validNums--
	}

	if isInteger(root2) {
		validNums--
	}
	return int(validNums)
}

func getQuadraticRoots(a, b, c float64) (float64, float64) {
	discriminantSqrt := math.Sqrt((b * b) - (4 * a * c))
	root1 := (-b + discriminantSqrt) / (2 * a)
	root2 := (-b - discriminantSqrt) / (2 * a)
	return root1, root2
}

func getAllNumsFromStr(str string) []int {
	numRe := regexp.MustCompile(`\d+`)
	numsAsStrs := numRe.FindAllString(str, -1)

	nums := []int{}
	for _, numAsStr := range numsAsStrs {
		num, _ := strconv.Atoi(numAsStr)
		nums = append(nums, num)
	}
	return nums
}

func extractAndCombineDigits(input string) int {
	re := regexp.MustCompile(`\d`)
	matches := re.FindAllString(input, -1)

	digitString := ""
	for _, match := range matches {
		digitString += match
	}

	number, _ := strconv.Atoi(digitString)
	return number
}

func isInteger(num float64) bool {
	return math.Mod(num, 1) == 0
}
