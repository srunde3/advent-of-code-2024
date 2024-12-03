package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CalculateDistance(left []int, right []int) int {
	if len(left) != len(right) {
		return 0
	}

	sort.Ints(left)
	sort.Ints(right)

	sum := 0

	for index, l := range left {
		r := right[index]
		diff := absDiff(l, r)
		sum += diff
	}

	return sum
}

func absDiff(left int, right int) int {
	result := left - right
	if result < 0 {
		return -result
	}
	return result
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var left []int
	var right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) == 2 {
			leftVal, err := strconv.Atoi((parts[0]))
			if err != nil {
				log.Fatalf("Failed to parse int from %s", parts[0])
			}
			rightVal, err := strconv.Atoi((parts[1]))
			if err != nil {
				log.Fatalf("Failed to parse int from %s", parts[1])
			}

			left = append(left, leftVal)
			right = append(right, rightVal)
		} else {
			log.Fatalf("Failed to parse data: %s", parts)
		}
	}

	distance := CalculateDistance(left, right)
	fmt.Printf("Distance: %d\n", distance)
}
