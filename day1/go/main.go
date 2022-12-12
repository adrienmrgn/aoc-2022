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
	var maxCalories = 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if calories > maxCalories {
				maxCalories = calories
			}
			calories = 0
		} else {
			intLine, _ := strconv.Atoi(line)
			calories += intLine
		}
	}
	fmt.Printf("Max calories is : %d", maxCalories)
}
