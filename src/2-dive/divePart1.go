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
    // Open file and read in one line at a time
    totalDistance, totalDepth = parseFile()

    // Output final total
    fmt.Println("Final distance:", totalDistance, "Final depth:", totalDepth)
    fmt.Println("Product:", totalDistance * totalDepth)
}

func parseFile() (int, int) {
    
    file, err := os.Open("input.txt")
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    distance := 0
    depth := 0

    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Fields(line)
        fmt.Println(line)
        fmt.Println("direction", fields[0], "change", fields[1])
        fmt.Println("")
        // Convert string value to integer
        intVal, err := strconv.Atoi(fields[1])
        check(err)
        
        distance, depth = applyCommand(distance, depth, fields[0], intVal)
    }
    
    check(scanner.Err())
    return distance, depth
}

func applyCommand(distance int, depth int, direction string, change int) (newDistance int, newDepth int) {
    switch direction {
    case "forward":
        return distance + change, depth
    case "down":
        return distance, depth + change
    case "up":
        return distance, depth - change
    default:
        log.Fatal()
    }
    return distance, depth
}
