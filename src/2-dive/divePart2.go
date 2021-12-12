package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

func check(e error) {
    if e != nil {
        log.Fatal()
    }
}

func main() {
    var totalDistance int
    var totalDepth int
    var finalAim int
    // Open file and read in one line at a time
    totalDistance, totalDepth, finalAim = parseFile()

    // Output final total
    fmt.Println("Final distance:", totalDistance, "Final depth:", totalDepth, "Final aim:", finalAim)
    fmt.Println("Product:", totalDistance * totalDepth)
}

func parseFile() (int, int, int) {
    
    file, err := os.Open("input.txt")
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    distance := 0
    depth := 0
    aim := 0

    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Fields(line)
        fmt.Println(line)
        fmt.Println("direction", fields[0], "change", fields[1])
        fmt.Println("")
        // Convert string value to integer
        intVal, err := strconv.Atoi(fields[1])
        check(err)
        
        distance, depth, aim = applyCommand(distance, depth, aim, fields[0], intVal)
    }
    
    check(scanner.Err())
    return distance, depth, aim
}

func applyCommand(distance int, depth int, aim int, direction string, change int) (newDistance int, newDepth int, newAim int) {
    switch direction {
    case "forward":
        return distance + change, depth + (aim * change), aim
    case "down":
        return distance, depth, aim + change
    case "up":
        return distance, depth, aim - change
    default:
        log.Fatal()
    }
    // This line should never be executed, but the compiler will complain without it
    return distance, depth, aim
}
