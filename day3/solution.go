package main

import (
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)


func part_1(data []byte) int {
    re := regexp.MustCompile(`mul\((\d+,\d+)\)`)
    expressions := re.FindAllSubmatch(data, -1)

    var total int = 0

    for _, expression := range expressions {
        if len(expression) != 2 {
            panic("Found an expression that had more than one submatch")
        }

        expression_s := strings.Split(string(expression[1]), ",")

        if len(expression_s) != 2 {
            panic("Found an expression that has more than two arguments")
        }

        lhs, err := strconv.Atoi(expression_s[0])

        if err != nil {
            panic(err)
        }

        rhs, err := strconv.Atoi(expression_s[1])

        if err != nil {
            panic(err)
        }

        total += lhs * rhs
    }

    return total
}


func part_2(data []byte) int {
    var total int = 0

    re := regexp.MustCompile(`do\(\)`)
    dos := re.Split(string(data), -1)

    for _, do := range dos {
        re := regexp.MustCompile(`don't\(\)`)
        donts := re.Split(do, -1)

        total += part_1([]byte(donts[0]))
    }

    return total
}


func main() {
    data, err := os.ReadFile("input.txt")

    if err != nil {
        panic(err)
    }

    part_1 := part_1(data)
    part_2 := part_2(data)

    fmt.Println(part_1)
    fmt.Println(part_2)
}
