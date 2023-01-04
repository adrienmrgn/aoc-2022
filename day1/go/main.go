package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input/day1.txt")
	defer file.Close()
	if err != nil {
		log.Fatalf("Failed to open file : %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var calories int

	// Will contain the top 3 values, orderd descendant
	topCalories := []int{0, 0, 0}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// Check if new calory count is greater that one of the top calory counters
			for i, value := range topCalories {
				if calories > value {
					// If a new calory count in greater than one of the current top, it wil take it place
					insertIntoSlice(topCalories, calories, i)
					// Breaking as if one enter the top, it's greater than the below ones
					break
				}
			}
			calories = 0
		} else {
			intLine, _ := strconv.Atoi(line)
			calories += intLine
		}
	}
	fmt.Printf("Top 3 Max calories is : \n")
	var sumOfTop int = 0
	for i, top := range topCalories {
		fmt.Printf("> %d : %d\n", i, top)
		sumOfTop += top
	}
	fmt.Printf("Sum of top 3 calory counters is : %v\n", sumOfTop)
}

func insertIntoSlice(slice []int, element int, index int) []int {
	// Store initial slice length as we want to keep the top a fixed size
	var initialLen = len(slice)
	// head is the slice composed of number left of the insertion point
	head := slice[:index]
	// tail is the slice composed of numbers right of the insertion point, minus 1 : the kicked-out number
	tail := slice[index : initialLen-1]
	// composed slice is head + new element + tail
	slice = append(head, append([]int{element}, tail...)...)
	return slice
}
