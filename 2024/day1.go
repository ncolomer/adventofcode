package adventofcode2024

import (
	"bufio"
	"fmt"
	"slices"
	"strings"
)

func Day1Read(input string) (list1, list2 []int, err error) {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		var e1, e2 int
		if _, err = fmt.Sscan(line, &e1, &e2); err != nil {
			err = fmt.Errorf("failed to parse line '%s': %v", line, err)
			return
		}
		list1 = append(list1, e1)
		list2 = append(list2, e2)
	}
	return
}

// SolveDay1Part1 compute total distance
func SolveDay1Part1(list1 []int, list2 []int) (res int, err error) {
	// store input slice in place
	slices.Sort(list1)
	slices.Sort(list2)
	// compute total distance
	for i, _ := range list1 {
		distance := list1[i] - list2[i]
		if distance < 0 {
			distance = -distance
		}
		res += distance
	}
	return
}

// SolveDay1Part2 compute similarity score
func SolveDay1Part2(list1 []int, list2 []int) (res int, err error) {
	// traverse right list and count value occurrences
	index := make(map[int]int)
	for _, x := range list2 {
		index[x] += 1
	}
	// compute similarity score
	for _, x := range list1 {
		res += x * index[x]
	}
	return
}
