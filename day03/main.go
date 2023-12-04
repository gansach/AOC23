package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/gansach/AOC23/pkg/fileutil"
)

func main() {
	lines, _ := fileutil.ReadFileLines("inputs/day03.txt")
	A(lines)
	B(lines)
}

func A(lines []string) {
	numsAndAdjacentSpecialChars := getNumsAndAdjacentSpecialChars(lines)
	sum := 0
	for _, numAndAdjacentSpecialChars := range numsAndAdjacentSpecialChars {
		if len(numAndAdjacentSpecialChars.adjacentSpecialChars) > 0 {
			sum += numAndAdjacentSpecialChars.num
		}
	}
	fmt.Println(sum)
}

func B(lines []string) {
	numsAndAdjacentSpecialChars := getNumsAndAdjacentSpecialChars(lines)
	starSpecialCharAdjacentNums := map[AdjacentSpecialChar][]int{}

	for _, numAndAdjacentSpecialChars := range numsAndAdjacentSpecialChars {
		num := numAndAdjacentSpecialChars.num
		adjacentSpecialChars := numAndAdjacentSpecialChars.adjacentSpecialChars
		for specialChar := range adjacentSpecialChars {
			if specialChar.ch == '*' {
				starSpecialCharAdjacentNums[specialChar] = append(starSpecialCharAdjacentNums[specialChar], num)
			}
		}
	}

	sum := 0
	for _, nums := range starSpecialCharAdjacentNums {
		if len(nums) == 2 {
			sum += nums[0] * nums[1]
		}
	}
	fmt.Println(sum)
}

type AdjacentSpecialChar struct {
	i  int
	j  int
	ch rune
}

type NumAndAdjacentSpecialChars struct {
	num                  int
	adjacentSpecialChars map[AdjacentSpecialChar]bool
}

func getNumsAndAdjacentSpecialChars(grid []string) []NumAndAdjacentSpecialChars {
	m := len(grid[0])
	numsAndAdjacentSpecialChars := []NumAndAdjacentSpecialChars{}

	for i, row := range grid {
		digitIndices := [][]int{}
		num := 0

		for j, char := range row {
			if unicode.IsDigit(char) {
				digit, _ := strconv.Atoi(string(char))
				num = num*10 + digit
				digitIndices = append(digitIndices, []int{i, j})
			}

			if !unicode.IsDigit(char) || j == m-1 {
				if num > 0 {
					adjacentSpecialChars := map[AdjacentSpecialChar]bool{}
					for _, pos := range digitIndices {
						for specialChar := range getAdjacentSpecialChars(pos[0], pos[1], grid) {
							adjacentSpecialChars[specialChar] = true
						}
					}

					numAndAdjacentSpecialChars := NumAndAdjacentSpecialChars{num, adjacentSpecialChars}
					numsAndAdjacentSpecialChars = append(numsAndAdjacentSpecialChars, numAndAdjacentSpecialChars)

					num = 0
					digitIndices = [][]int{}
				}
			}
		}
	}

	return numsAndAdjacentSpecialChars
}

func checkListHasAdjacentSpecialChar(indices [][]int, grid []string) bool {
	for _, pos := range indices {
		adjacentSpecialChars := getAdjacentSpecialChars(pos[0], pos[1], grid)
		if len(adjacentSpecialChars) > 0 {
			return true
		}
	}
	return false
}

func getAdjacentSpecialChars(i, j int, grid []string) map[AdjacentSpecialChar]bool {
	n := len(grid)
	m := len(grid[0])
	adjacentSpecialChars := map[AdjacentSpecialChar]bool{}

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	for _, direction := range directions {
		ni := i + direction[0]
		nj := j + direction[1]

		if isValid(ni, nj, n, m) && isSpecialChar(rune(grid[ni][nj])) {
			adjacentSpecialChar := AdjacentSpecialChar{ni, nj, rune(grid[ni][nj])}
			adjacentSpecialChars[adjacentSpecialChar] = true
		}
	}
	return adjacentSpecialChars
}

func isValid(i, j, n, m int) bool {
	return i >= 0 && i < n && j >= 0 && j < m
}

func isSpecialChar(ch rune) bool {
	return !unicode.IsDigit(ch) && ch != '.'
}
