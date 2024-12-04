package main

import (
    "fmt"
    "os"
    "strings"
)


func has_target(
    data []string,
    target string,
    row int,
    col int,
    row_delta int,
    col_delta int,
) bool {
    for i := range len(target) {
        if data[row+i*row_delta][col+i*col_delta] != target[i] {
            return false
        }
    }

    return true
}


func part_1(data []string) int {
    // Find all horizontal, vertical, and diagonal instances of XMAS
    var n_rows = len(data)
    var target string = "XMAS"
    var n_target int = len(target)
    var n_appearances int = 0

    for row := range len(data) {
        var n_cols int = len(data[row])

        for col := range n_cols {
            var north bool = row >= n_target - 1
            var east bool = col <= n_cols - n_target
            var south bool = row <= n_rows - n_target
            var west bool = col >= n_target - 1

            if north && has_target(data, target, row, col, -1, 0) {
                n_appearances += 1
            }

            if north && east && has_target(data, target, row, col, -1, 1) {
                n_appearances += 1
            }

            if east && has_target(data, target, row, col, 0, 1) {
                n_appearances += 1
            }

            if south && east && has_target(data, target, row, col, 1, 1) {
                n_appearances += 1
            }

            if south && has_target(data, target, row, col, 1, 0) {
                n_appearances += 1
            }

            if south && west && has_target(data, target, row, col, 1, -1) {
                n_appearances += 1
            }

            if west && has_target(data, target, row, col, 0, -1) {
                n_appearances += 1
            }

            if north && west && has_target(data, target, row, col, -1, -1) {
                n_appearances += 1
            }
        }
    }

    return n_appearances
}


func part_2(data []string) int {
    // Find all instances of MAS in the shape of an X
    var n_rows int = len(data)
    var target string = "MAS"
    var radius int = len(target) / 2
    var n_appearances int = 0

    for row := 1; row < n_rows - radius; row++ {
        var n_cols int = len(data[row])

        for col := 1; col < n_cols - radius; col++ {
            northeast := has_target(data, target, row+1, col-1, -1, 1)
            southeast := has_target(data, target, row-1, col-1, 1, 1)
            southwest := has_target(data, target, row-1, col+1, 1, -1)
            northwest := has_target(data, target, row+1, col+1, -1, -1)

            if (northeast || southwest) && (southeast || northwest) {
                n_appearances += 1
            }
        }
    }

    return n_appearances
}


func main() {
    data, err := os.ReadFile("input.txt")

    if err != nil {
        panic(err)
    }

    // Ignore empty rows
    var rows []string

    for _, row := range strings.Split(string(data), "\n") {
        if row != "" {
            rows = append(rows, row)
        }
    }

    part_1 := part_1(rows)
    part_2 := part_2(rows)

    fmt.Println(part_1)
    fmt.Println(part_2)
}
