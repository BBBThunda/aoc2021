package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
//    "strings"
)

func check(e error) {
    if e != nil {
        log.Fatal()
    }
}

func main() {
    var count0 [12]int
    var count1 [12]int
    var gammaRate, epsilonRate string
    var gammaRateDec, epsilonRateDec int64
    
    // Open file and read in one line at a time
    parseFile(&count0, &count1)

    for index, _ := range count0 {
        if count0[index] > count1[index] {
            gammaRate += "0"
            epsilonRate += "1"
        } else if count1[index] > count0[index] {
            gammaRate += "1"
            epsilonRate += "0"
        } else {
            log.Fatal("A tie was detected, ABORT!")
        }
    }

    gammaRateDec, _ = strconv.ParseInt(gammaRate, 2, 64)
    epsilonRateDec, _ = strconv.ParseInt(epsilonRate, 2, 64)

    // Output final total
    fmt.Println("Gamma Rate:", gammaRate, "Epsilon Rate:", epsilonRate)
    fmt.Println("Gamma Rate:", gammaRateDec, "Epsilon Rate:", epsilonRateDec)
    fmt.Println("Product:", gammaRateDec * epsilonRateDec)
}

func parseFile(count0 *[12]int, count1 *[12]int) {
    file, err := os.Open("input.txt")
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        processLine(scanner.Text(), count0, count1)
    }

    check(scanner.Err())
}

func processLine(line string, count0 *[12]int, count1 *[12]int) {
    for index, character := range line {
        if string(character) == "0" {
            (*count0)[index]++
        } else if string(character) == "1" {
            (*count1)[index]++
        } else {
            log.Fatal()
        }
    }
}
