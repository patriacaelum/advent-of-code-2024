// Solution: 2 815 556
package main

import (
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)


func main() {
    data, err := os.ReadFile("input.txt")

    if err != nil {
        panic(err)
    }

    rows := strings.Split(string(data), "\n")

    var left_col []int
    var right_col []int
    var n_rows int = 0

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
        n_rows += 1
    }

    sort.Ints(left_col)
    sort.Ints(right_col)

    var total_distance int = 0

    for i := 0; i < n_rows; i++ {
        var distance int = left_col[i] - right_col[i]

        if distance < 0 {
            distance *= -1
        }

        total_distance += distance
    }

    fmt.Println(total_distance)
}
