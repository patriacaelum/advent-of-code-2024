package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)


func sign(x int) int {
    if x > 0 {
        return 1
    } else if x < 0 {
        return -1
    }

    return 0
}


func is_safe(report []int) bool {
    // Checks if each report follows one of two rules:
    // 1. Each consecutive number is all increasing or decreasing
    // 2. Any two adjacent number differs by at least 1 and at most 3

    var diff_min int = 1
    var diff_max int = 3

    var direction int = 0
    var prior int = 0

    for index, level := range report {
        if index == 0 {
            prior = level
            continue
        }

        var diff int = level - prior

        if index == 1 {
            // Set the direction on first diff
            direction = sign(diff)

            if direction == 0 {
                return false
            }
        }

        // Check first condition
        if sign(diff) != direction {
            return false
        }

        // Check second condition
        diff *= direction

        if diff < diff_min || diff > diff_max {
            return false
        }

        prior = level
    }

    return true
}


func part_1(reports [][]int) int {
    var n_safe int = 0

    for _, report := range reports {
        if is_safe(report) {
            n_safe += 1
        }
    }

    return n_safe
}


func part_2(reports [][]int) int {
    var n_safe int = 0

    for _, report := range reports {
        if is_safe(report) {
            n_safe += 1
            continue
        }

        for i := 0; i < len(report); i++ {
            var subreport []int

            subreport = append(subreport, report[:i]...)

            if i + 1 < len(report) {
                subreport = append(subreport, report[i + 1:]...)
            }

            if is_safe(subreport) {
                n_safe += 1
                break
            }
        }
    }

    return n_safe
}


func main() {
    data, err := os.ReadFile("input.txt")

    if err != nil {
        panic(err)
    }

    var reports [][]int

    for _, report_data := range strings.Split(string(data), "\n") {
        var report []int

        if len(report_data) == 0 {
            continue
        }

        for _, i := range strings.Fields(report_data) {
            value, err := strconv.Atoi(i)

            if err != nil {
                panic(err)
            }

            report = append(report, value)
        }

        reports = append(reports, report)
    }

    part_1 := part_1(reports)
    part_2 := part_2(reports)

    fmt.Println(part_1)
    fmt.Println(part_2)
}
