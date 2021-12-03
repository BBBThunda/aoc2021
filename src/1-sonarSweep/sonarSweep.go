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
	var increases, previousValue int
	firstLine := true
	for scanner.Scan() {
		check(scanner.Err())
		text := scanner.Text()
		currentValue, scErr := strconv.Atoi(text)
		check(scErr)
		fmt.Println("Text: ", text, "Current: ", currentValue, "Previous: ", previousValue)
		if !firstLine && currentValue > previousValue {
			increases++
			fmt.Println("INCREASE")
		}
		previousValue = currentValue
		firstLine = false
	}
	fmt.Println("Number of increases: ", increases)
}
