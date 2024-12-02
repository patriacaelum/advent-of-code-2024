// Solution: 2 815 556
package main

import (
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)


func load_rows(rows []string) ([]int, []int) {
    var left_col []int
    var right_col []int

    for _, row := range rows {
        items := strings.Fields(row)

        if len(items) != 2 {
            continue
        }

        left_item, err := strconv.Atoi(items[0])

        if err != nil {
            panic(err)
        }

        right_item, err := strconv.Atoi(items[1])

        if err != nil {
            panic(err)
        }

        left_col = append(left_col, left_item)
        right_col = append(right_col, right_item)
    }

    return left_col, right_col
}


func part_1(left_col []int, right_col []int) int {
    // Sort columns and calculate distance between sorted columns
    sort.Ints(left_col)
    sort.Ints(right_col)

    var total_distance int = 0
    var n_rows int = len(left_col)

    for i := 0; i < n_rows; i++ {
        var distance int = left_col[i] - right_col[i]

        if distance < 0 {
            distance *= -1
        }

        total_distance += distance
    }

    return total_distance
}


func part_2(left_col []int, right_col []int) int {
    // Calculate the similarity score between the columns, equal to the number
    // of times a number in the left column appears in the right column
    
    // Find the number of times a number appears in the right column
    factor := make(map[int]int)

    for _, value := range right_col {
        existing := factor[value]
        factor[value] = existing + 1
    }

    // Calculate the similarity score
    var similarity int = 0

    for _, value := range left_col {
        n, ok := factor[value]

        if ok {
            similarity += value * n
        }
    }

    return similarity
}


func main() {
    data, err := os.ReadFile("input.txt")

    if err != nil {
        panic(err)
    }

    rows := strings.Split(string(data), "\n")
    left_col, right_col := load_rows(rows)

    part_1 := part_1(left_col, right_col)
    part_2 := part_2(left_col, right_col)

    fmt.Println(part_1)
    fmt.Println(part_2)
}
