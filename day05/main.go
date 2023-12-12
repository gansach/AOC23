package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	mathAOC "github.com/gansach/AOC23/pkg/math"
)

type Node struct {
	Value  string
	Ranges [][]int
	Next   *Node
}

func main() {
	content, _ := os.ReadFile("inputs/day05.txt")
	contentChunks := strings.Split(strings.TrimSpace(string(content)), "\r\n\r\n")

	seeds := getAllNumsFromStr(contentChunks[0])
	var head *Node = makeLinkedList(contentChunks[1:])

	fmt.Println(A(seeds, head))
	B(seeds, head)
}

func A(seeds []int, head *Node) int {
	minTransformation := math.MaxInt
	for _, seed := range seeds {
		transformedVal := applyTransformationsToList(seed, head)
		minTransformation = mathAOC.Min(minTransformation, transformedVal)
	}
	return minTransformation
}

func B(seedRanges []int, head *Node) {
	minTransformation := math.MaxInt

	for i := 0; i < len(seedRanges); i += 2 {
		seedRangeStart := seedRanges[i]
		seedRangeLen := seedRanges[i+1]

		var seeds []int
		for i := 0; i < seedRangeLen; i++ {
			seeds = append(seeds, seedRangeStart+i)
		}

		minTransformation = mathAOC.Min(minTransformation, A(seeds, head))
		fmt.Println(minTransformation)
	}
  fmt.Println("Part B: ", minTransformation)
}

func makeLinkedList(mappings []string) *Node {
	var head *Node = nil
	current := head

	for _, chunk := range mappings {
		mapping := strings.Split(chunk, "\r\n")

		sourceRe := regexp.MustCompile(`(\w+)-to-\w+`)
		sourceMatches := sourceRe.FindStringSubmatch(strings.TrimSpace(mapping[0]))

		source := sourceMatches[1]

		rangesArr := [][]int{}
		for _, mappingRange := range mapping[1:] {
			rangesArr = append(rangesArr, getAllNumsFromStr(mappingRange))
		}

		node := Node{source, rangesArr, nil}

		if head == nil {
			head = &node
			current = head
		} else {
			current.Next = &node
			current = current.Next
		}
	}
	return head
}

func applyTransformationsToList(initialVal int, head *Node) int {
	val := initialVal
	for head != nil {
		val = transform(val, head.Ranges)
		head = head.Next
	}
	return val
}

func transform(source int, mappings [][]int) int {
	for _, mapping := range mappings {
		destinationRangeStart := mapping[0]
		sourceRangeStart := mapping[1]
		rangeLen := mapping[2]

		if source >= sourceRangeStart && source < (sourceRangeStart+rangeLen) {
			return destinationRangeStart + (source - sourceRangeStart)
		}
	}
	return source
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
