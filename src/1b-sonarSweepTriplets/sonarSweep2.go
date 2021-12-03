package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("HOLY CRAP I CAN GO!")
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var increases, penultimate, antepenultimate, previousTriplet int
	firstLineRead, secondLineRead, thirdLineRead := false, false, false

	for scanner.Scan() {
		var currentTriplet int
		check(scanner.Err())
		text := scanner.Text()
		ultimate, scErr := strconv.Atoi(text)
		check(scErr)

		// Calculate current values
		if secondLineRead {
			currentTriplet = ultimate + penultimate + antepenultimate
		}
		fmt.Println("Text: ", text, " Current: ", currentTriplet, " Previous: ", previousTriplet)
		if thirdLineRead {
			if currentTriplet > previousTriplet {
				increases++
				fmt.Println("INCREASE")
			}
		}
		// Retain values for next iteration
		antepenultimate = penultimate
		penultimate = ultimate
		previousTriplet = currentTriplet
		if secondLineRead {
			thirdLineRead = true
		}
		if firstLineRead {
			secondLineRead = true
		}
		firstLineRead = true
	}
	fmt.Println("Number of increases: ", increases)
}
