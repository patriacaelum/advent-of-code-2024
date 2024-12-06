package main

import (
    "fmt"
    "os"
    "slices"
    "strconv"
    "strings"
)


func update_is_correct(requirements map[string][]string, update []string) bool {
    for i := 0; i < len(update) - 1; i++ {
        for _, requirement := range requirements[update[i]] {
            if slices.Contains(update[i + 1:], requirement) {
                return false
            }
        }
    }

    return true
}


func find_median(
    requirements map[string][]string,
    update []string,
    k int,
) string {
    // Clever solution I stole from reddit.
    // Since part 2 only asks for the median value, we don't actually need to
    // sort the misordered array
    // We can just check the number of elements that *must* be prior to a
    // proposed element, and return it if it turns out to be the median,
    // otherwise, make another guess
    median_index := len(update) / 2
    proposal := update[k]

    if len(update) == 1 {
        return proposal
    }

    var lhs []string

    for _, precedent := range requirements[proposal] {
        if slices.Contains(update, precedent) {
            lhs = append(lhs, precedent)
        }
    }

    fmt.Println(lhs)

    if len(lhs) == median_index {
        return proposal
    }

    return find_median(requirements, update, k + 1)
}


func part_1_and_2(requirements []string, updates []string) (int, int) {
    // Build a map of preceding required updates
    requires := make(map[string][]string)

    for _, requirement := range requirements {
        cols := strings.Split(requirement, "|")

        if len(cols) != 2 {
            panic("Requirement has incorrect number of columns!")
        }

        before := cols[0]
        after := cols[1]

        _, ok := requires[after]

        if !ok {
            requires[after] = []string{before}
        } else {
            requires[after] = append(requires[after], before)
        }
    }

    // Check each update to see if they follow the required order
    var correct_sum int = 0
    var incorrect_sum int = 0

    for _, update := range updates {
        if update == "" {
            continue
        }

        pages := strings.Split(update, ",")

        // Then take center page and sum them up
        if update_is_correct(requires, pages) {
            page_no, err := strconv.Atoi(pages[len(pages) / 2])

            if err != nil {
                panic(err)
            }

            correct_sum += page_no
        } else {
            fmt.Println("Pages", pages)
            median := find_median(requires, pages, 0)
            page_no, err := strconv.Atoi(median)

            if err != nil {
                panic(err)
            }

            incorrect_sum += page_no
        }
    }

    return correct_sum, incorrect_sum
}


func main() {
    data, err := os.ReadFile("input.txt")

    if err != nil {
        panic(err)
    }

    sections := strings.Split(string(data), "\n\n")

    if len(sections) < 2 {
        panic("Sections not parsed correctly!")
    }

    requirements := strings.Split(sections[0], "\n")
    updates := strings.Split(sections[1], "\n")

    part_1, part_2 := part_1_and_2(requirements, updates)

    fmt.Println(part_1)
    fmt.Println(part_2)
}
