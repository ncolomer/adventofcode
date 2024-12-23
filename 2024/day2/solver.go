package day2

import (
	"bufio"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Day2Read(input string) (reports [][]int, err error) {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	re := regexp.MustCompile(`\d+`)
	for scanner.Scan() {
		line := scanner.Text()
		var levels []int
		for _, match := range re.FindAllString(line, -1) {
			value, _ := strconv.Atoi(match)
			levels = append(levels, value)
		}
		reports = append(reports, levels)
	}
	return
}

// SolveDay2Part1 compute number of safe report
func SolveDay2Part1(reports [][]int) (res int, err error) {
	for _, levels := range reports {
		if isReportSafe(levels) {
			res += 1
		}
	}
	return res, nil
}

// SolveDay2Part2 compute number of safe report with Problem Dampener
func SolveDay2Part2(reports [][]int) (res int, err error) {
out:
	for _, levels := range reports {
		// first attempt, using raw report
		if isReportSafe(levels) {
			res += 1
			continue
		}
		// second attempt, try to pop each index separately until we find a safe report
		for i := 0; i < len(levels); i++ {
			head := levels[:i]
			tail := levels[i+1:]
			cleaned := slices.Concat(head, tail)
			if isReportSafe(cleaned) {
				res += 1
				// we found one safe report, continue to next report
				continue out
			}
		}
	}
	return res, nil
}

func isReportSafe(levels []int) bool {
	sign := levels[1] < levels[0]
	for i := 0; i < len(levels)-1; i++ {
		a, b := levels[i], levels[i+1]
		absDiff := b - a
		if absDiff < 0 {
			absDiff = -absDiff
		}
		if b == a {
			return false
		}
		if (b < a) != sign {
			return false
		}

		if absDiff > 3 {
			return false
		}
	}
	return true
}
