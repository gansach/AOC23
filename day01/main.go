package main

import (
	"fmt"

	"github.com/gansach/AOC23/pkg/fileutil"
	"github.com/gansach/AOC23/pkg/math"
)

func main() {
	lines, _ := fileutil.ReadFileLines("inputs/day01.txt")
	A(lines)
	B(lines)
}

func A(lines []string) {
	digitMap := make(map[string]int)
	for i := 1; i <= 9; i++ {
		key := fmt.Sprintf("%d", i)
		digitMap[key] = i
	}
	sum := findCalibrationValueSum(lines, digitMap)
	fmt.Println(sum)
}

func B(lines []string) {
	digitMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for i := 1; i <= 9; i++ {
		key := fmt.Sprintf("%d", i)
		digitMap[key] = i
	}
	sum := findCalibrationValueSum(lines, digitMap)
	fmt.Println(sum)
}

func findCalibrationValueSum(lines []string, digitMap map[string]int) int {
	sum := 0
	for _, line := range lines {
		firstDigitIdxMap := createDigitIdxMap()
		lastDigitIdxMap := createDigitIdxMap()

		for digitStr, digit := range digitMap {
			firstIndex, lastIndex := getFirstAndLastOccurence(line, digitStr)

			if firstIndex != -1 {
				if firstDigitIdxMap[digit] != -1 {
					firstDigitIdxMap[digit] = math.Min(firstDigitIdxMap[digit], firstIndex)
				} else {
					firstDigitIdxMap[digit] = firstIndex
				}
			}

			if lastIndex != -1 {
				if lastDigitIdxMap[digit] != -1 {
					lastDigitIdxMap[digit] = math.Max(lastDigitIdxMap[digit], lastIndex)
				} else {
					lastDigitIdxMap[digit] = lastIndex
				}
			}
		}

		firstDigit := reduceDigitIdxMap(firstDigitIdxMap, math.Min)
		lastDigit := reduceDigitIdxMap(lastDigitIdxMap, math.Max)

		sum += firstDigit*10 + lastDigit
	}
	return sum
}

func reduceDigitIdxMap(digitIndexMap map[int]int, findExtreme func(int, int) int) int {
	extremeIndex := -1
	extremeIndexDigit := -1
	for digit, index := range digitIndexMap {
		if index == -1 {
			continue
		}

		if extremeIndex == -1 {
			extremeIndex = index
			extremeIndexDigit = digit
		} else {
			extremeIndex = findExtreme(extremeIndex, index)
			if extremeIndex == index {
				extremeIndexDigit = digit
			}
		}
	}
	return extremeIndexDigit
}

func createDigitIdxMap() map[int]int {
	digitMap := make(map[int]int)
	for i := 1; i <= 9; i++ {
		digitMap[i] = -1
	}
	return digitMap
}

// Can use KMP
func getFirstAndLastOccurence(str, digit string) (int, int) {
	digitSize := len(digit)
	strSize := len(str)

	firstIndex := -1
	lastIndex := -1

	for i := 0; i < strSize; i++ {
		if i+digitSize <= strSize && str[i] == digit[0] && digit == str[i:i+digitSize] {
			if firstIndex == -1 {
				firstIndex = i
			}
			lastIndex = i
		}
	}

	return firstIndex, lastIndex
}
